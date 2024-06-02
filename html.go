package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func generateErrorNode(err error, msg string) elem.Node {
	return elem.Body(nil,
		elem.H1(nil,
			elem.Text(msg),
		),
		elem.P(nil, elem.Text(err.Error())),
	)
}

func generateHeadNode(title string, description string) elem.Node {
	return elem.Head(nil,
		elem.Title(nil, elem.Text(title)),
		elem.Meta(attrs.Props{attrs.Charset: "utf-8"}),
		elem.Meta(attrs.Props{attrs.Name: "description", attrs.Content: description}),
		elem.Meta(attrs.Props{attrs.Name: "viewport", attrs.Content: "width=device-width, initial-scale=1"}),
	)
}

func GenerateHomeHTML() string {
	head := generateHeadNode("Home", "Home Description")

	body := elem.Div(nil,
		elem.H1(nil, elem.Text("Test Homepage")),
		elem.A(attrs.Props{attrs.Href: "/recipes"}, elem.Text("View All Recipes")),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}

func GenerateRecipesHTML() string {
	head := generateHeadNode("Recipes", "View All Recipes")

	recs, err := GetAllRecipes()
	if err != nil {
		body := generateErrorNode(err, "Unable to load recipes; please try again later.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}
	sort.Slice(recs, func(i, j int) bool {
		return recs[i].Name < recs[j].Name
	})

	body := elem.Body(nil)
	for _, rec := range recs {
		body.Children = append(body.Children,
			elem.H1(nil,
				elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/recipes/%v", rec.ID)},
					elem.Text(rec.Name)),
			),
		)
	}

	html := elem.Html(nil, head, body)

	return html.Render()
}

func GenerateRecipeDetailHTML(id string) string {
	recID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		head := elem.Head(nil, elem.Title(nil, elem.Text("Error loading recipe")))
		body := generateErrorNode(err, "Malformed recipe ID.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}

	rec, err := GetRecipeByID(recID)
	if err != nil {
		head := elem.Head(nil, elem.Title(nil, elem.Text("Error loading recipe")))
		body := generateErrorNode(err, "Unable to load recipe.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}

	head := generateHeadNode(rec.Name, rec.Description)

	ings := elem.Ul(nil)
	for _, ing := range rec.Ingredients {
		ings.Children = append(ings.Children,
			elem.Li(nil, elem.Text(fmt.Sprintf("%v %v %v", ing.Quantity, ing.Unit, ing.Label))),
		)
	}

	instr := elem.Ol(nil)
	for _, step := range rec.Instructions {
		instr.Children = append(instr.Children,
			elem.Li(nil, elem.Text(step)),
		)
	}

	body := elem.Body(nil,
		elem.Header(nil, elem.H1(nil, elem.Text(rec.Name))),
		elem.P(nil, elem.Text(rec.Description)),
		elem.Main(nil,
			elem.H2(nil, elem.Text("Ingredients")),
			ings,
			elem.H2(nil, elem.Text("Instructions")),
			instr,
		),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}
