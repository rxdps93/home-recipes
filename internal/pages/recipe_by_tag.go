package pages

import (
	"fmt"
	"sort"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/rxdps93/home-recipes/internal/db"
)

func GenerateRecipesByTagHTML(tag string) string {
	head := GenerateHeadNode(fmt.Sprintf("Recipes By Tag: %v", tag), fmt.Sprintf("Recipes Tagged With %v", tag))

	recs, err := db.GetAllRecipesForTagName(tag)
	if err != nil {
		head := GenerateHeadNode("Error", "Unable to load recipes")
		body := GenerateErrorNode(err, "Unable to load recipes")
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

	body := GenerateBodyStructure("Recipes By Tag",
		elem.H1(nil, elem.Text(fmt.Sprintf("Recipes Tagged With %v", tag))),
		GenerateJumpLinks(ltrs),
		GenerateJumpDestinations(ltrs, sections),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}
