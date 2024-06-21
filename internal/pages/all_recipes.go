package pages

import (
	"fmt"
	"sort"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/rxdps93/home-recipes/internal/db"
)

func recipeJumpLinks(ltrs []rune) elem.Node {
	content := elem.Div(attrs.Props{attrs.Class: "rec-jump"},
		elem.Hr(nil),
		elem.H3(nil, elem.Text("Jump to Section...")),
	)

	for i, ltr := range ltrs {
		content.Children = append(content.Children,
			elem.A(attrs.Props{attrs.Href: "#" + string(ltr)}, elem.Text(string(ltr))),
		)

		if i != len(ltrs)-1 {
			content.Children = append(content.Children, elem.Text("&middot;"))
		} else {
			content.Children = append(content.Children, elem.Hr(nil))
		}
	}

	return content
}

func recipeJumpDestinations(ltrs []rune, sections map[rune][]elem.Node) elem.Node {
	content := elem.Div(attrs.Props{attrs.Class: "rec-links"})

	for _, ltr := range ltrs {
		content.Children = append(content.Children,
			elem.Div(attrs.Props{attrs.Class: "rec-sec", attrs.ID: string(ltr)},
				elem.H3(nil, elem.Text(string(ltr))),
				elem.Ul(nil,
					sections[ltr]...,
				),
			),
		)
	}

	return content
}

func GenerateRecipesHTML() string {
	head := GenerateHeadNode("Recipes", "View All Recipes")

	recs, err := db.GetAllRecipes()
	if err != nil {
		body := GenerateErrorNode(err, "Unable to load recipes; please try again later.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}

	sections := make(map[rune][]elem.Node)
	for _, rec := range recs {
		sections[rune(rec.Name[0])] = append(sections[rune(rec.Name[0])],
			elem.Li(nil,
				elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/recipes/%v", rec.ID)},
					elem.Text(rec.Name),
				),
			),
		)
	}

	ltrs := make([]rune, 0)
	for k := range sections {
		ltrs = append(ltrs, k)
	}
	sort.Slice(ltrs, func(i, j int) bool {
		return ltrs[i] < ltrs[j]
	})

	body := GenerateBodyStructure("Family Recipe Index",
		elem.H1(nil, elem.Text("View All Recipes")),
		recipeJumpLinks(ltrs),
		recipeJumpDestinations(ltrs, sections),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}
