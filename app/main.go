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
	fmt.Println("serving at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
