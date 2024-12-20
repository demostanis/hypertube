package components

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/demostanis/hypertube/mvdb"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var CategoryIndex = 0

var ScrollLeftAttr = `event.preventDefault();
		const container = document.getElementById('%s');
		const RightArrow = document.getElementById('%s-right');
		const LeftArrow = document.getElementById('%s-left');

		const opacity = (container.scrollLeft - container.offsetWidth <= 0) ? 0 : 1;
		LeftArrow.style.zIndex = opacity;
		RightArrow.style.zIndex = 1;
		LeftArrow.style.opacity = opacity;
		RightArrow.style.opacity = 1;
		
		container.scrollLeft -= container.offsetWidth;
		return false;`

var ScrollRightAttr = `event.preventDefault();
		const container = document.getElementById('%s');
		const RightArrow = document.getElementById('%s-right');
		const LeftArrow = document.getElementById('%s-left');
		const maxScrollLeft = container.scrollWidth - container.clientWidth;
		
		const opacity = (container.scrollLeft + container.offsetWidth >= maxScrollLeft) ? 0 : 1;
		RightArrow.style.zIndex = opacity;
		LeftArrow.style.zIndex = 1;
		RightArrow.style.opacity = opacity;
		LeftArrow.style.opacity = 1;
		
		container.scrollLeft += container.offsetWidth;
		return false;`

func Card(poster string) Node {
	return Div(Class("card"),
		Div(Class("card-image"),
			Figure(Class("image is-4by5"),
				Img(
					Class("poster-file"),
					Attr("style", "height: 121%;"),
					Src("https://image.tmdb.org/t/p/w500/"+poster),
				),
			),
		),
	)
}

func ScrollArrows(categoryId string) Node {
	return Div(Class("control arrows is-hidden-mobile"),
		Button(
			Class("arrow-left"),
			ID(fmt.Sprintf("%s-left", categoryId)),
			Attr("onclick", fmt.Sprintf(ScrollLeftAttr, categoryId, categoryId, categoryId)),
			Span(Class("icon"),
				I(Class("fa-solid fa-chevron-left"), Attr("style", "scale: 1.5;")),
			),
		),
		Button(
			Class("arrow-right"),
			ID(fmt.Sprintf("%s-right", categoryId)),
			Attr("onclick", fmt.Sprintf(ScrollRightAttr, categoryId, categoryId, categoryId)),
			Span(Class("icon"),
				I(Class("fa-solid fa-chevron-right"), Attr("style", "scale: 1.5;")),
			),
		),
	)
}

func CreateCards(ContentList mvdb.ApiResponse, categoryId string) []Node {
	CardArray := make([]Node, len(ContentList.Results))

	for i, Content := range ContentList.Results {
		Title := Content.Title
		if Title == "" {
			Title = Content.Name
		}

		Content.Overview = strings.ReplaceAll(Content.Overview, "\"", "")
		CardArray[i] = Div(Class("column pl-0 pr-5"), ID(categoryId+"-"+strconv.Itoa(i)),
			Attr("style", "display: flex;"),
			Div(
				Class("cell is-clickable"),
				Attr("hx-get", "/show-content-popup"),
				Attr("hx-trigger", "click"),
				Attr("hx-target", "#content-popup"),
				Attr("hx-swap", "innerHTML"),
				Attr("hx-vals", fmt.Sprintf(`{"Id": %d, "isMovie": "%t"}`,
					Content.Id, (Content.Title != ""))),
				Card(Content.PosterPath),
			),
		)
	}
	return CardArray
}

func CreateCardGrill(ContentList mvdb.ApiResponse, categoryId string) Node {
	CardArray := CreateCards(ContentList, categoryId)

	return Div(Class("list"),
		ScrollArrows(categoryId),
		Div(
			append([]Node{
				Class("columns is-mobile pl-5"),
				ID(categoryId),
				Attr(
					"style",
					"overflow-x: auto; "+
						"flex-wrap: nowrap; "+
						"margin: 0; "+
						"scroll-behavior: smooth; "+
						"position: relative;",
				),
			}, CardArray[:]...)...,
		),
	)
}

func CreateCategory(ContentList mvdb.ApiResponse, Name string) Node {
	categoryId := fmt.Sprintf("category-grid-%d", CategoryIndex)
	CategoryIndex = CategoryIndex + 1
	return Div(
		Class("category"),
		Div(
			Class("category-title title is-2 ml-5 mt-3 mb-1"),
			Attr("style", "position: relative;"),
			Text(Name),
		),
		CreateCardGrill(ContentList, categoryId),
	)
}

func ContentCategory(Request, CategoryName string) Node {
	var ContentList mvdb.ApiResponse

	json.Unmarshal(mvdb.CallMvdbDefault(Request), &ContentList)

	Category := CreateCategory(ContentList, CategoryName)

	if CategoryIndex == 99 {
		CategoryIndex = 0
	}

	return Category
}
