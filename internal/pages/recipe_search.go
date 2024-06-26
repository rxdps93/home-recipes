package pages

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/rxdps93/home-recipes/internal/db"
)

// TODO: this ought to utilize HTMX
func GenerateRecipeTable(recs []db.Recipe) elem.Node {
	return nil
}

func GenerateRecipeSearchHTML() string {
	head := GenerateHeadNode("Search Recipes", "Search Recipes")

	body := GenerateBodyStructure("Search Recipes",
		GenerateRecipeNavLinks(),
		elem.H2(attrs.Props{attrs.Style: "text-align: center"}, elem.Text("Placeholder for search box + filters")),
		elem.H2(attrs.Props{attrs.Style: "text-align: center"}, elem.Text("Placeholder for search results")),
	)

	return elem.Html(nil, head, body).Render()
}
