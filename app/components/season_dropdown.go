package components

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var disableDropdown = `event.preventDefault();
		const checkbox = document.getElementById('dropdown-toggle');
		const dropdownText = document.getElementById('dropdown-text');
		const dropdownElement = document.getElementById('%d');
		checkbox.checked = false;
		dropdownText.textContent = dropdownElement.textContent;
		return false;`

func genDropDownContent(items []string, contentID int) Node {
	dropDownItems := make([]Node, len(items))
	var special int

	if items[0] == "Épisodes spéciaux" {
		special = 0
	} else {
		special = 1
	}

	for i, item := range items {
		dropDownItems[i] = A(
			Class("dropdown-item"),
			ID(fmt.Sprintf("%d", i)),
			Text(item),
			Attr("onclick", fmt.Sprintf(disableDropdown, i)),
			Attr("hx-get", "/switch-season"),
			Attr("hx-trigger", "click"),
			Attr("hx-target", "#episode-grid-container"),
			Attr("hx-swap", "innerHTML"),
			Attr("hx-vals", fmt.Sprintf(`{"Id": %d, "season_num": "%d"}`,
				contentID, i+special)),
		)
	}
	dropDownContent := append([]Node{Class("dropdown-content")},
		dropDownItems...)
	return Div(dropDownContent...)
}

func DropDown(items []string, contentID int) Node {
	return Div(
		Class("dropdown"),
		Attr("style", "left: 6.4vw; margin-top: 1em;"),
		Input(
			ID("dropdown-toggle"),
			Attr("type", "checkbox"),
		),
		Div(
			Class("dropdown-trigger"),
			Button(
				Class("button"),
				ID("dropdown-button"),
				Span(ID("dropdown-text"), Text("Saison 1")),
				Span(Class("icon is-small"), I(Class("fas fa-angle-down"))),
			),
		),
		Div(
			Class("dropdown-menu"),
			ID("dropdown-menu"),
			Attr("role", "menu"),
			genDropDownContent(items, contentID),
		),
	)
}
