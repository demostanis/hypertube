package pages

import (
	"net/http"

	"github.com/demostanis/hypertube/components"
	"github.com/demostanis/hypertube/mvdb"
	"github.com/demostanis/hypertube/pages"
	. "maragu.dev/gomponents"
)

func handleEmpty(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleShowFilmPopup(w http.ResponseWriter, r *http.Request) {
	FilmId := r.URL.Query().Get("filmId")
	FilmTitle := r.URL.Query().Get("titlefilm")
	FilmOverview := r.URL.Query().Get("overview")
	FilmImage := r.URL.Query().Get("image")
	TrailerLink := mvdb.FindTrailer(FilmId)

	FilmCard := components.CreatePopup(FilmTitle, FilmOverview, TrailerLink, FilmImage)

	w.Header().Set("Content-Type", "text/html")
	err := FilmCard.Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
