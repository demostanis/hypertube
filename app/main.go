package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/pat"

	ghttp "maragu.dev/gomponents/http"
)

func main() {
	r := pat.New()

	r.Get("/login", ghttp.Adapt(LoginHandler))
	r.Post("/login", APILoginHandler)
	r.Get("/", ghttp.Adapt(HomeHandler))

	http.Handle("/", r)
	port, ok := os.LookupEnv("port")
	if !ok {
		port = "8080"
	}
	fmt.Println("serving at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}