package pages

import (
	"net/http"

	"github.com/demostanis/hypertube/components"
	"github.com/demostanis/hypertube/mvdb"
	. "maragu.dev/gomponents"
)

func HandleEmpty(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HandleShowContentPopup(w http.ResponseWriter, r *http.Request) {
	ContentId := r.URL.Query().Get("Id")
	ContentTitle := r.URL.Query().Get("title")
	ContentOverview := r.URL.Query().Get("overview")
	ContentImage := r.URL.Query().Get("image")
	TrailerLink := mvdb.FindTrailer(ContentId)

	ContentCard := components.CreatePopup(ContentTitle, ContentOverview, TrailerLink, ContentImage)

	w.Header().Set("Content-Type", "text/html")
	err := ContentCard.Render(w)
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

func InternalErrorHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return InternalError(), nil
}
