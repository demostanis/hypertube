package pages

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/demostanis/hypertube/components"
	"github.com/demostanis/hypertube/mvdb"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func SerieContent(Serie mvdb.Content) Node {
	var SeasonList mvdb.Content
	var EpisodeList mvdb.Content

	json.Unmarshal(mvdb.CallMvdbDefault("https://api.themoviedb.org/3/tv/"+strconv.Itoa(Serie.Id)+"?append_to_response=episode_groups&language=fr-FR"), &SeasonList)
	json.Unmarshal(mvdb.CallMvdbDefault("https://api.themoviedb.org/3/tv/"+strconv.Itoa(Serie.Id)+"/season/1?language=fr-FR"), &EpisodeList)

	SeasonNames := make([]string, len(SeasonList.Seasons))

	for i, Season := range SeasonList.Seasons {
		SeasonNames[i] = Season.Name
	}

	return Div(
		components.DropDown(SeasonNames, Serie.Id),
		Div(
			ID("episode-grid-container"),
			components.EpisodeGrid(EpisodeList.Episodes, Serie.ImagePath),
		),
	)
}

func ContentPage(contentType, Query string) Node {
	var contentList mvdb.ApiResponse
	var PageContent Node
	var Name string

	json.Unmarshal(mvdb.CallMvdbDefault("https://api.themoviedb.org/3/search/"+contentType+"?query="+url.QueryEscape(Query)+"&include_adult=true&language=fr-FR&page=1"), &contentList)
	fmt.Println(contentList)
	if len(contentList.Results) == 0 {
		fmt.Println("oui")
		return Home()
	}
	if contentType == "movie" {
		Name = contentList.Results[0].Title
		PageContent = nil
	} else if contentType == "tv" {
		Name = contentList.Results[0].Name
		PageContent = SerieContent(contentList.Results[0])
	}

	return components.Page(
		Class("has-navbar-fixed-top"),
		components.Navbar(),
		Img(
			Class("is-hidden-touch"),
			Src("https://image.tmdb.org/t/p/original/"+contentList.Results[0].ImagePath),
			Attr("style", "filter: blur(5px); z-index: -1; position: fixed; width: 100%; opacity: 0.5;"),
		),
		Img(
			Class("is-hidden-desktop"),
			Src("https://image.tmdb.org/t/p/original/"+contentList.Results[0].PosterPath),
			Attr("style", "filter: blur(5px); z-index: -1; position: fixed; height: 100%; opacity: 0.5;"),
		),
		Div(
			ID("player-container"),
			Attr("style", "display: flex; justify-content: center; margin-top: 9vh;"),
			components.Player(Name, contentList.Results[0].ImagePath, contentList.Results[0].Overview, contentList.Results[0].Id, (contentType == "movie"), false),
		),
		PageContent,
		components.ContentCategory(fmt.Sprintf("https://api.themoviedb.org/3/%s/%d/recommendations?language=fr-FR&page=1", contentType, contentList.Results[0].Id), "Recommendations"),
		components.Foot())
}
