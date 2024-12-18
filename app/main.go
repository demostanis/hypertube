//go:nofmt
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/demostanis/hypertube/api"
	"github.com/demostanis/hypertube/models"
	"github.com/demostanis/hypertube/pages"
	"github.com/gorilla/pat"

	ghttp "maragu.dev/gomponents/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/empty", pages.HandleEmpty)
	http.HandleFunc("/show-content-popup", pages.HandleShowContentPopup)

	r := pat.New()

	r.Get("/login", ghttp.Adapt(pages.LoginHandler))
	r.Post("/login", api.APILoginHandler)
	r.Get("/signin", ghttp.Adapt(pages.SigninHandler))
	r.Post("/signin", api.APISigninHandler)
	r.Get("/", ghttp.Adapt(pages.HomeHandler))

	http.Handle("/", r)
	port, ok := os.LookupEnv("port")
	if !ok {
		port = "8080"
	}

	db, err := models.ConnectToDatabase(
		"crocotube",
		"crocotube",
		os.Getenv("HYPERTUBE_DB_PASSWORD"),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to connect db:", err)
		return
	}
	fmt.Println("Connected to db:", db)

	test := models.Content{ID: 1, Name: "The Great DinoMalin"}
	fmt.Println(test.Name)

	fmt.Println("serving at http://0.0.0.0:" + port)
	http.ListenAndServe(":"+port, nil)
}
