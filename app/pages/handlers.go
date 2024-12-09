package pages

import (
	"fmt"
	"net/http"
	"strconv"

	. "maragu.dev/gomponents"
)

func handleScrollLeft(w http.ResponseWriter, r *http.Request) {
	scrollLeft := r.URL.Query().Get("scrollLeft")
	offsetWidth := r.URL.Query().Get("offsetWidth")

	scrollLeftInt, _ := strconv.Atoi(scrollLeft)
	offsetWidthInt, _ := strconv.Atoi(offsetWidth)

	if scrollLeftInt-offsetWidthInt <= 0 {
		fmt.Fprintf(w, "0")
	} else {
		fmt.Fprintf(w, "1")
	}
}

func handleScrollRight(w http.ResponseWriter, r *http.Request) {
	scrollLeft := r.URL.Query().Get("scrollLeft")
	offsetWidth := r.URL.Query().Get("offsetWidth")
	maxScroll := r.URL.Query().Get("maxScroll")

	scrollLeftInt, _ := strconv.Atoi(scrollLeft)
	offsetWidthInt, _ := strconv.Atoi(offsetWidth)
	maxScrollInt, _ := strconv.Atoi(maxScroll)

	if scrollLeftInt+offsetWidthInt >= maxScrollInt {
		fmt.Fprintf(w, "0")
	} else {
		fmt.Fprintf(w, "1")
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return Home(), nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return Login(), nil
}

func SigninHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return Signin(), nil
}
