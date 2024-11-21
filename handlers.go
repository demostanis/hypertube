package main

import (
	"net/http"

	"github.com/demostanis/hypertube/pages"
	. "maragu.dev/gomponents"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return pages.Home(), nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return pages.Login(), nil
}
