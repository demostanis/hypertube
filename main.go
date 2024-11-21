package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	ghttp "maragu.dev/gomponents/http"
)

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
