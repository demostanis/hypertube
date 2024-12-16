//go:nofmt
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/demostanis/hypertube/api"
	"github.com/demostanis/hypertube/pages"
	"github.com/gorilla/pat"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	ghttp "maragu.dev/gomponents/http"
)

func connectToDatabase() (*gorm.DB, error) {
	pass := os.Getenv("HYPERTUBE_DB_PASSWORD")
	dsn := fmt.Sprintf("host=postgres user=crocotube password=%s dbname=crocotube port=5432", pass)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

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

	db, err := connectToDatabase()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to connect db:", err)
		return
	}
	fmt.Println("Connected to db:", db)

	fmt.Println("serving at http://0.0.0.0:" + port)
	http.ListenAndServe(":"+port, nil)
}
