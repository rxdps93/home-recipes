package pages

import (
	"fmt"
	"sort"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/rxdps93/home-recipes/internal/db"
)

func GenerateRecipesHTML() string {
	head := GenerateHeadNode("Recipes", "View All Recipes", false)

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
				elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/recipes/%v", rec.ID), attrs.Class: "link"},
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

	body := GenerateBodyStructure("View All Recipes",
		GenerateRecipeNavLinks(),
		GenerateJumpLinks(ltrs),
		GenerateJumpDestinations(ltrs, sections),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}
