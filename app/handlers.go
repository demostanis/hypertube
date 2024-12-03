package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/demostanis/hypertube/pages"
	. "maragu.dev/gomponents"
)

func handleScrollLeft(w http.ResponseWriter, r *http.Request) {
	scrollLeft := r.URL.Query().Get("scrollLeft")

	scrollPos, _ := strconv.Atoi(scrollLeft)

	if scrollPos-1883 <= 0 {
		fmt.Fprintf(w, "0")
	} else {
		fmt.Fprintf(w, "1")
	}
}

func handleScrollRight(w http.ResponseWriter, r *http.Request) {
	scrollLeft := r.URL.Query().Get("scrollLeft")

	scrollPos, _ := strconv.Atoi(scrollLeft)

	if scrollPos+1883 >= 1883*2 {
		fmt.Fprintf(w, "0")
	} else {
		fmt.Fprintf(w, "1")
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return pages.Home(), nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return pages.Login(), nil
}
