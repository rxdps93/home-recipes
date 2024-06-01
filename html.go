package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GenerateHomeHTML() string {
	head := elem.Head(nil, elem.Title(nil, elem.Text("Home")))

	body := elem.Div(nil,
		elem.H1(nil, elem.Text("Test Homepage")),
		elem.A(attrs.Props{attrs.Href: "/recipes"}, elem.Text("View All Recipes")),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}

func GenerateRecipesHTML() string {
	head := elem.Head(nil, elem.Title(nil, elem.Text("Recipes")))

	recs, err := GetAllRecipes()
	if err != nil {
		// TODO: handle this error case properly
		log.Fatal(err)
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
		// TODO: handle this error case properly
		log.Fatal(err)
	}

	rec, err := GetRecipeByID(recID)
	if err != nil {
		// TODO: handle this error case properly
		log.Fatal(err)
	}

	head := elem.Head(nil, elem.Title(nil, elem.Text(rec.Name)))

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
