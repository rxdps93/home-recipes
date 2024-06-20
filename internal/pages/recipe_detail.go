package pages

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/rxdps93/home-recipes/internal/db"
)

// TODO: update html structure & styling
func GenerateRecipeDetailHTML(id string) string {
	recID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		head := GenerateHeadNode("Error", "Unable to Load Recipe")
		body := GenerateErrorNode(err, "Malformed recipe ID.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}

	rec, err := db.GetRecipeByID(recID)
	if err != nil {
		head := GenerateHeadNode("Error", "Unable to Load Recipe")
		body := GenerateErrorNode(err, "Unable to load recipe.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}

	head := GenerateHeadNode(rec.Name, rec.Description)

	ings := elem.Ul(nil,
		elem.TransformEach(rec.Ingredients, func(ing db.Ingredient) elem.Node {
			return elem.Li(nil, elem.Text(fmt.Sprintf("%v %v %v", ing.Quantity, ing.Unit, ing.Label)))
		})...,
	)

	instr := elem.Ol(nil,
		elem.TransformEach(rec.Instructions, func(step string) elem.Node {
			return elem.Li(nil, elem.Text(step))
		})...,
	)

	tags := elem.Ul(nil,
		elem.TransformEach(rec.Tags, func(tag string) elem.Node {
			return elem.Li(nil,
				elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/tags/%v", tag)},
					elem.Text(tag)),
			)
		})...,
	)

	var src elem.Node
	_, err = url.ParseRequestURI(rec.Source)
	if err != nil {
		src = elem.Text(fmt.Sprintf(rec.Source))
	} else {
		src = elem.A(attrs.Props{attrs.Href: rec.Source}, elem.Text(rec.Source))
	}

	body := elem.Body(nil,
		GenerateNavigationHTML(),
		elem.Header(nil, elem.H1(nil, elem.Text(rec.Name))),
		elem.P(nil, elem.Text(rec.Description)),
		elem.Main(nil,
			elem.H2(nil, elem.Text("Ingredients")),
			ings,
			elem.H2(nil, elem.Text("Instructions")),
			instr,
			elem.H3(nil, elem.Text("Tags")),
			tags,
			elem.B(nil, elem.Text("Recipe Source: ")),
			src,
		),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}
