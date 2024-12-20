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
	contentId := r.URL.Query().Get("Id")
	isMovie := r.URL.Query().Get("isMovie")
	trailerLink := mvdb.GetTrailer(contentId, (isMovie == "true"))
	content := mvdb.GetContentByID(contentId, "fr-FR", (isMovie == "true"))
	contentCard := components.CreatePopup(content.Title, content.Overview, trailerLink, content.ImagePath)

	w.Header().Set("Content-Type", "text/html")
	err := contentCard.Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return Home(), nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return Login(""), nil
}

func SigninHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return Signin(""), nil
}

func InternalErrorHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	return InternalError(), nil
}
