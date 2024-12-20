package pages

import (
	"encoding/json"
	"net/http"

	"github.com/demostanis/hypertube/components"
	"github.com/demostanis/hypertube/mvdb"
	. "maragu.dev/gomponents"
)

func HandleSelectEpisode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	thumbnail := r.URL.Query().Get("thumbnail")
	overview := r.URL.Query().Get("overview")

	player := components.Player(name, thumbnail, overview, -1, false, true)
	w.Header().Set("Content-Type", "text/html")
	err := player.Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HanldeSwitchSeason(w http.ResponseWriter, r *http.Request) {
	contentId := r.URL.Query().Get("Id")
	seasonNum := r.URL.Query().Get("season_num")
	var content mvdb.Content
	var season mvdb.Content

	json.Unmarshal(mvdb.CallMvdbDefault("https://api.themoviedb.org/3/tv/"+contentId+"?language=fr-FR"), &content)
	json.Unmarshal(mvdb.CallMvdbDefault("https://api.themoviedb.org/3/tv/"+contentId+"/season/"+seasonNum+"?language=fr-FR"), &season)

	episodeGrid := components.EpisodeGrid(season.Episodes, content.ImagePath)
	w.Header().Set("Content-Type", "text/html")
	err := episodeGrid.Render(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleEmpty(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HandleShowContentPopup(w http.ResponseWriter, r *http.Request) {
	contentId := r.URL.Query().Get("Id")
	isMovie := r.URL.Query().Get("isMovie")
	trailerLink := mvdb.GetTrailer(contentId, (isMovie == "true"))
	content := mvdb.GetContentByID(contentId, "fr-FR", (isMovie == "true"))
	contentCard := components.CreatePopup(
		content.Title, content.Overview, trailerLink, content.ImagePath)

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

func ContentHandler(w http.ResponseWriter, r *http.Request) (Node, error) {
	contentType := r.URL.Query().Get(":type")

	if contentType != "movie" && contentType != "tv" {
		return Home(), nil
	}

	return ContentPage(contentType, r.URL.Query().Get(":query")), nil
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
