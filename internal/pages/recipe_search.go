package pages

import (
	"log"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
	"github.com/rxdps93/home-recipes/internal/db"
)

func GenerateTableBody(nameQuery string, tagQuery []string) string {
	recs, err := db.GetRecipesFiltered(nameQuery, tagQuery)
	if err != nil {
		log.Printf("GenerateTableBody: %v\n", err)
		return elem.Text("An error has occurred").Render()
	}

	if len(recs) == 0 {
		return elem.Text("No recipes match the given criteria").Render()
	}

	var content string
	for _, rec := range recs {
		tagElements := elem.TransformEach(rec.Tags, func(tag string) elem.Node {
			return elem.Li(nil, elem.Text(tag))
		})

		content += elem.Tr(nil,
			elem.Td(attrs.Props{attrs.Class: "sr-name"}, elem.Text(rec.Name)),
			elem.Td(attrs.Props{attrs.Class: "sr-tags"},
				elem.Ul(nil,
					tagElements...,
				),
			),
		).Render()
	}

	return content
}

func generateTagDropdown() elem.Node {
	dropdown := elem.Select(attrs.Props{
		attrs.Name:     "tags",
		attrs.ID:       "tags",
		attrs.Class:    "tag-filter",
		attrs.Multiple: "true",
		htmx.HXPost:    "/search",
		htmx.HXTrigger: "load, change delay:500mx, tags",
		htmx.HXTarget:  "#search-results",
		htmx.HXInclude: "[name='name']",
	})

	tags, err := db.GetAllTags()
	if err != nil {
		log.Printf("generateTagDropdown: %v\n", err)
		return dropdown
	}

	for _, tag := range tags {
		dropdown.Children = append(dropdown.Children, elem.Option(attrs.Props{attrs.Value: tag}, elem.Text(tag)))
	}

	return dropdown
}

func GenerateRecipeSearchHTML() string {
	head := GenerateHeadNode("Search Recipes", "Search Recipes", true)

	body := GenerateBodyStructure("Search Recipes",
		GenerateRecipeNavLinks(),
		elem.Div(attrs.Props{attrs.Class: "rec-search"},
			elem.Div(attrs.Props{attrs.Class: "rec-filters"},
				elem.Input(attrs.Props{
					attrs.Class:       "name-filter",
					attrs.Type:        "search",
					attrs.Name:        "name",
					attrs.Placeholder: "Search by Recipe Name...",
					htmx.HXPost:       "/search",
					htmx.HXTrigger:    "load, input changed delay:500ms, name",
					htmx.HXTarget:     "#search-results",
					htmx.HXInclude:    "[name='tags']",
				}),
				elem.Label(nil, elem.Text("Tags")),
				generateTagDropdown(),
			),
			elem.Hr(nil),
			elem.Table(attrs.Props{attrs.Class: "rec-table"},
				elem.THead(nil,
					elem.Tr(nil,
						elem.Th(nil, elem.Text("Name")),
						elem.Th(nil, elem.Text("Tags")),
					),
				),
				elem.TBody(attrs.Props{attrs.ID: "search-results"}),
			),
		),
	)

	return elem.Html(nil, head, body).Render()
}
