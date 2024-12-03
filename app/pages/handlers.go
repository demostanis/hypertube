package pages

import (
	"fmt"
	"net/http"
	"strconv"

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
	return Home(), nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return Login(), nil
}

func SigninHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return Signin(), nil
}
