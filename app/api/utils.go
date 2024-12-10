package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/schema"
	. "maragu.dev/gomponents"
	ghttp "maragu.dev/gomponents/http"
)

var decoder = schema.NewDecoder()

func bad(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request"))
}

type Handler = func(string) Node

func apiError(
	w http.ResponseWriter,
	r *http.Request,
	h Handler,
	err string,
) (Node, error) {
	wrapper := func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return h(err), nil
	}
	ghttp.Adapt(wrapper)(w, r)
	return wrapper(w, r)
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

	link := "http://keycloak:8000/realms/" + realm + "/protocol/openid-connect/token"
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
		return "", errors.New(data.Error)
	}

	return data.AccessToken, nil
}

func adminAuthorization() (string, error) {
	return auth("admin-cli",
		os.Getenv("KEYCLOAK_ADMIN"),
		os.Getenv("KEYCLOAK_ADMIN_PASSWORD"),
		"master")
}
