//go:nofmt
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/demostanis/hypertube/api"
	"github.com/demostanis/hypertube/pages"
	"github.com/gorilla/pat"

	ghttp "maragu.dev/gomponents/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/empty", pages.HandleEmpty)
	http.HandleFunc("/show-content-popup", pages.HandleShowContentPopup)
	http.HandleFunc("/switch-season", pages.HanldeSwitchSeason)
	http.HandleFunc("/select-episode", pages.HandleSelectEpisode)

	r := pat.New()

	r.Get("/login", ghttp.Adapt(pages.LoginHandler))
	r.Post("/login", api.APILoginHandler)
	r.Get("/signin", ghttp.Adapt(pages.SigninHandler))
	r.Post("/signin", api.APISigninHandler)
	r.Get("/content/{type}/{query}", ghttp.Adapt(pages.ContentHandler))
	r.Get("/", ghttp.Adapt(pages.HomeHandler))

	http.Handle("/", r)
	port, ok := os.LookupEnv("port")
	if !ok {
		port = "8080"
	}
	fmt.Println("serving at http://0.0.0.0:" + port)
	http.ListenAndServe(":"+port, nil)
}
