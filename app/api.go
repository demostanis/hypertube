package main

import (
	"net/http"

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

func APILoginHandler(w http.ResponseWriter, r *http.Request) {
	var params LoginParams

	paramsInto(&params, w, r)

	if params.Username == "admin" && params.Password == "admin" {
		// do stuff I suppose
	}

	ghttp.Adapt(LoginHandler)(w, r)
}
