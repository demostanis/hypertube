package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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

// API endpoints

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	AccessToken string `json:"access_token"`
	Error       string `json:"error"`
}

// connect to a user in a realm
func APILoginHandler(w http.ResponseWriter, r *http.Request) {
	var params LoginParams

	paramsInto(&params, w, r)

	form := url.Values{}
	form.Add("client_id", "hypertube-auth")
	form.Add("username", params.Username)
	form.Add("password", params.Password)
	form.Add("grant_type", "password")

	link := "http://keycloak:8080/realms/default/protocol/openid-connect/token"
	res, err := http.PostForm(link, form)
	if err != nil {
		fmt.Printf("Error logging: %s\n", err.Error())
		return
	}
	defer res.Body.Close()

	var data Response
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Printf("Could not read error: %s", err.Error())
		return
	}

	if data.Error != "" {
		fmt.Printf("error: %s\n", data.Error)
	} else {
		fmt.Printf("access_token: %s\n", data.AccessToken)
	}

	ghttp.Adapt(LoginHandler)(w, r)
}

func APISigninHandler(w http.ResponseWriter, r *http.Request) {
	ghttp.Adapt(SigninHandler)(w, r)
}
