package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/demostanis/hypertube/pages"
	"github.com/gorilla/mux"

	. "maragu.dev/gomponents"

	ghttp "maragu.dev/gomponents/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return pages.Home(), nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return pages.Login(), nil
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", ghttp.Adapt(HomeHandler))
	r.HandleFunc("/login", ghttp.Adapt(LoginHandler))

	http.Handle("/", r)
	port, ok := os.LookupEnv("port")
	if !ok {
		port = "8080"
	}
	fmt.Println("serving at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
