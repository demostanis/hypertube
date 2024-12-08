//go:nofmt
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/pat"

	ghttp "maragu.dev/gomponents/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/empty", handleEmpty)
	http.HandleFunc("/show-content-popup", handleShowContentPopup)

	r := pat.New()

	r.Get("/login", ghttp.Adapt(LoginHandler))
	r.Post("/login", APILoginHandler)
	r.Get("/", ghttp.Adapt(HomeHandler))

	http.Handle("/", r)
	port, ok := os.LookupEnv("port")
	if !ok {
		port = "8080"
	}
	fmt.Println("serving at http://0.0.0.0:" + port)
	http.ListenAndServe(":"+port, nil)
}
