package pages

import (
	"fmt"
	"log"
	"sort"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
	"github.com/rxdps93/home-recipes/internal/db"
)

func GenerateTestTableBody(nameQuery string, tagQuery []string) string {
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

func generateTestMultiSelect() elem.Node {
	tags, err := db.GetAllTags()
	if err != nil {
		log.Printf("generateTagDropdown: %v\n", err)
		return elem.Text("Unable to load tags")
	}

	sort.Strings(tags)

	fieldset := elem.Fieldset(attrs.Props{
		attrs.Name:     "tags",
		attrs.ID:       "tags",
		attrs.Class:    "tag-filter",
		htmx.HXPost:    "/search-test",
		htmx.HXTrigger: "load, change delay:500ms, tags",
		htmx.HXTarget:  "#search-results",
		htmx.HXInclude: "[name='name'],[name='tags']",
	}, elem.Legend(nil, elem.Text("Filter By Tag")))

	list := elem.Ul(attrs.Props{attrs.Class: "tag-filter-list"})

	for _, tag := range tags {
		list.Children = append(list.Children, elem.Li(attrs.Props{attrs.Class: "link"},
			elem.Input(attrs.Props{
				attrs.Type:  "checkbox",
				attrs.Name:  "tags",
				attrs.ID:    fmt.Sprintf("chkbox-%v", tag),
				attrs.Value: tag,
			}),
			elem.Label(attrs.Props{
				attrs.For: fmt.Sprintf("chkbox-%v", tag),
			}, elem.Text(tag)),
		))
	}

	fieldset.Children = append(fieldset.Children, list)

	return fieldset
}

func GenerateTestHTML() string {
	head := GenerateHeadNode("Test Page", "Test", true)

	body := GenerateBodyStructure("Test Page",
		GenerateRecipeNavLinks(),
		elem.Div(attrs.Props{attrs.Class: "rec-search"},
			elem.Div(attrs.Props{attrs.Class: "rec-filters"},
				elem.Fieldset(attrs.Props{attrs.Class: "name-filter-box"},
					elem.Legend(nil, elem.Text("Filter By Name")),
					elem.Input(attrs.Props{
						attrs.Class:       "name-filter",
						attrs.Type:        "search",
						attrs.Name:        "name",
						attrs.Placeholder: "Search by Recipe Name...",
						htmx.HXPost:       "/search-test",
						htmx.HXTrigger:    "load, input changed delay:500ms, name",
						htmx.HXTarget:     "#search-results",
						htmx.HXInclude:    "[name='tags']",
					}),
				),
				generateTestMultiSelect(),
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
