package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func CreatePopupHeader(ContentTitle string) Node {
	ContentTitleMobile := ContentTitle
	if len(ContentTitle) >= 40 {
		ContentTitle = ContentTitle[:40] + "..."
	}
	if len(ContentTitleMobile) >= 24 {
		ContentTitleMobile = ContentTitleMobile[:24] + "..."
	}
	return Header(Class("modal-card-head"),
		P(
			Class("modal-card-title is-hidden-mobile"),
			Attr("style", "text-align: center; padding-left: 20px;"),
			Text(ContentTitle),
		),
		P(
			Class("modal-card-title is-hidden-tablet"),
			Attr("style", "text-align: center; padding-left: 20px;"),
			Text(ContentTitleMobile),
		),
		Button(Class("delete"),
			Attr("hx-get", "/empty"),
			Attr("hx-target", "closest .modal"),
			Attr("hx-swap", "delete"),
			Attr("aria-label", "close"),
		),
	)
}

func CreatePopupTrailer(TrailerLink, ContentImage string) Node {
	Style := "background-color: rgba(20, 22, 26, 1); padding: 15px; padding-bottom: 0px;"

	if TrailerLink != "" {
		return Div(
			Attr("style", "background-color: rgba(20, 22, 26, 1);"),
			IFrame(
				Class("is-hidden-mobile"),
				Src(TrailerLink),
				Attr("allowfullscreen"),
				Attr("frameborder", "0"),
				Attr("width", "640"),
				Attr("height", "360"),
				Attr("style", Style),
			),
			IFrame(
				Class("is-hidden-tablet"),
				Src(TrailerLink),
				Attr("allowfullscreen"),
				Attr("frameborder", "0"),
				Attr("style", "width: 100vw; height: calc(100vw / 1.77777777777);"),
			),
		)
	}
	if ContentImage != "" {
		return Img(
			Class("image is-16by9"),
			Src("https://image.tmdb.org/t/p/original"+ContentImage),
			Attr("style", Style),
		)
	}
	return Div(
		Text("No trailer was found about this content :("),
		Br(),
		Text("But don't worry! You can still watch it!"),
		Attr("style", Style+"text-align: center;"),
	)
}

func CreatePopupFooter() Node {
	return Footer(Class("modal-card-foot"),
		Div(Class("buttons"),
			Button(Class("button"), Attr("style", "outline: solid 1px mediumslateblue"),
				Span(Attr("style", "color: mediumslateblue;"), Text("PLAY")),
				Span(Class("icon"), Attr("style", "color: mediumslateblue;"),
					I(Class("fa-solid fa-play"), Attr("aria-hidden", "true")),
				),
			),
		),
	)
}

func CreatePopup(ContentTitle, ContentOverview, TrailerLink, ContentImage string) Node {
	return Div(Class("modal is-active"),
		Div(Class("modal-background")),
		Div(Class("modal-card"),
			CreatePopupHeader(ContentTitle),
			CreatePopupTrailer(TrailerLink, ContentImage),
			Section(Class("modal-card-body"),
				Text(ContentOverview),
			),
			CreatePopupFooter(),
		),
	)
}
