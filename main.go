package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
	ghttp "maragu.dev/gomponents/http"
)

func NavbarItem(name string, path string) Node {
	return A(Class("navbar-item"), Href(path), Text(name))
}

func Navbar() Node {
	return Nav(Class("navbar"),
		Div(Class("navbar-menu"),
			Div(Class("navbar-start"),
				NavbarItem("Home", "/"),
				NavbarItem("Videos", "/videos"),
			),
			Div(Class("navbar-end"),
				NavbarItem("Login", "/login"),
			),
		),
	)
}

func Home() Node {
	return Navbar()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return HTML5(HTML5Props{
		Title: "Hypertube",
		Head: []Node{
			Link(
				Rel("stylesheet"),
				Href("https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css"),
			),
		},
		Body: []Node{Home()},
	}), nil
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ghttp.Adapt(HomeHandler))

	http.Handle("/", r)
	port, ok := os.LookupEnv("port")
	if !ok {
		port = "8080"
	}
	fmt.Println("serving at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
