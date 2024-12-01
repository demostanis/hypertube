package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"

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

// connect to a user in a realm
// todo : get authorization with http://keycloak.localhost:8000/realms/default/protocol/openid-connect/auth
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
	body, err := ioutil.ReadAll(res.Body)	

	if err != nil {
		fmt.Printf("Could not read error: %s", err.Error())
		return
	}
	
	fmt.Println(string(body))

	ghttp.Adapt(LoginHandler)(w, r)
}
