package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/schema"
	ghttp "maragu.dev/gomponents/http"
)

// API utils

var decoder = schema.NewDecoder()

func bad(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request"))
}

func paramsInto(
	dst interface{},
	w http.ResponseWriter,
	r *http.Request,
) error {
	err := r.ParseForm()
	if err != nil {
		bad(w)
		return err
	}

	err = decoder.Decode(dst, r.PostForm)
	if err != nil {
		bad(w)
		return err
	}

	return nil
}

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SigninParams struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	PasswordCheck string `json:"passwordCheck"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
}

type Response struct {
	AccessToken string `json:"access_token"`
	Error       string `json:"error_description"`
}

func auth(cid string, username string, password string, realm string) (string, error) {
	form := url.Values{}
	form.Add("client_id", cid)
	form.Add("username", username)
	form.Add("password", password)
	form.Add("grant_type", "password")

	link := "http://keycloak:8080/realms/" + realm + "/protocol/openid-connect/token"
	res, err := http.PostForm(link, form)
	if err != nil {
		return "", fmt.Errorf("Error posting form: %w", err)
	}
	defer res.Body.Close()

	var data Response
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", fmt.Errorf("Error decoding response: %w", err)
	}

	if data.Error != "" {
		return "", fmt.Errorf("Error from keycloak: %s", data.Error)
	}

	return data.AccessToken, nil
}

func adminAuthorization() (string, error) {
	return auth("admin-cli",
		os.Getenv("KEYCLOAK_ADMIN"),
		os.Getenv("KEYCLOAK_ADMIN_PASSWORD"),
		"master")
}

func APILoginHandler(w http.ResponseWriter, r *http.Request) {
	var params LoginParams

	paramsInto(&params, w, r)

	token, err := auth("hypertube-auth", params.Username, params.Password, "default")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("access token: %s\n", token)
	}

	ghttp.Adapt(LoginHandler)(w, r)
}

func APISigninHandler(w http.ResponseWriter, r *http.Request) {
	var params SigninParams

	paramsInto(&params, w, r)

	token, err := adminAuthorization()
	if err != nil {
		fmt.Printf("Failed to get authorization: %w", err)
		return
	}

	fmt.Printf("access token: %s\n", token)

	if params.Password != params.PasswordCheck {
		fmt.Printf("Error: Password doesn't match\n")
		ghttp.Adapt(SigninHandler)(w, r)
		return
	}

	payload := map[string]interface{}{
		"username":  params.Username,
		"email":     params.Email,
		"enabled":   true,
		"firstName": params.FirstName,
		"lastName":  params.LastName,
		"credentials": []map[string]interface{}{
			{
				"type":      "password",
				"value":     params.Password,
				"temporary": false,
			},
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error serializing json: %w", err)
		return
	}

	req, err := http.NewRequest("POST",
		"http://keycloak:8080/admin/realms/default/users",
		bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error creating request: %w", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %w\n", err)
		return
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response: %w\n", err)
		return
	}

	fmt.Printf("User %s created\n", params.Username)

	ghttp.Adapt(SigninHandler)(w, r)
}
