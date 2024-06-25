package pages

import "github.com/chasefleming/elem-go"

func GenerateRecipeSearchHTML() string {
	head := GenerateHeadNode("Search Recipes", "Search Recipes")

	body := GenerateBodyStructure("Search Recipes", elem.Text("Search Recipes"))

	return elem.Html(nil, head, body).Render()
}
