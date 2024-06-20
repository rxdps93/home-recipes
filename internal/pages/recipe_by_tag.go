package pages

import (
	"fmt"
	"sort"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/rxdps93/home-recipes/internal/db"
)

// TODO: update html structure & styling
func GenerateRecipesByTagHTML(tag string) string {
	recs, err := db.GetAllRecipesForTagName(tag)
	if err != nil {
		head := GenerateHeadNode("Error", "Unable to load recipes")
		body := GenerateErrorNode(err, "Unable to load recipes")
		html := elem.Html(nil, head, body)
		return html.Render()
	}

	head := GenerateHeadNode(fmt.Sprintf("Recipes By Tag: %v", tag), fmt.Sprintf("Recipes Tagged As %v", tag))

	sort.Slice(recs, func(i, j int) bool {
		return recs[i].Name < recs[j].Name
	})

	tags := elem.Ul(nil,
		elem.TransformEach(recs, func(rec db.Recipe) elem.Node {
			return elem.Li(nil,
				elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/recipes/%v", rec.ID)},
					elem.Text(rec.Name),
				),
			)
		})...,
	)

	body := elem.Body(nil,
		GenerateNavigationHTML(),
		elem.H2(nil, elem.Text(fmt.Sprintf("Recipes Tagged As %v:", tag))),
		tags,
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}
