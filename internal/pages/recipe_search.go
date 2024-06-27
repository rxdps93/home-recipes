package pages

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
)

func GenerateTableBody() string {
	content := elem.Tr(nil,
		elem.Td(nil, elem.Text("Recipe 1")),
		elem.Td(nil, elem.Text("Family Recipe")),
		elem.Td(nil, elem.Text("Cookbook")),
	).Render()

	content += elem.Tr(nil,
		elem.Td(nil, elem.Text("Another Recipe")),
		elem.Td(nil, elem.Text("Mexican")),
		elem.Td(nil, elem.Text("based.cooking")),
	).Render()

	return content
}

func GenerateRecipeSearchHTML() string {
	head := GenerateHeadNode("Search Recipes", "Search Recipes", true)

	body := GenerateBodyStructure("Search Recipes",
		GenerateRecipeNavLinks(),
		elem.H3(nil,
			elem.Span(attrs.Props{attrs.Class: "htmx-indicator"},
				elem.Img(attrs.Props{attrs.Src: "/assets/imgs/bars.svg"}),
				elem.Text("Searching..."),
			),
		),
		elem.Input(attrs.Props{
			attrs.Class:       "form-control",
			attrs.Type:        "search",
			attrs.Name:        "search",
			attrs.Placeholder: "Begin Typing To Search Recipes...",
			htmx.HXPost:       "/search",
			htmx.HXTrigger:    "input changed delay:500ms, search",
			htmx.HXTarget:     "#search-results",
			htmx.HXIndicator:  ".htmx-indicator",
		}),

		elem.Table(attrs.Props{attrs.Class: "table"},
			elem.THead(nil,
				elem.Tr(nil,
					elem.Th(nil, elem.Text("Name")),
					elem.Th(nil, elem.Text("Tags")),
					elem.Th(nil, elem.Text("Source")),
				),
			),
			elem.TBody(attrs.Props{attrs.ID: "search-results"}),
		),
	)

	return elem.Html(nil, head, body).Render()
}
