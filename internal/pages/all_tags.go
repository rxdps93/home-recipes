package pages

import (
	"fmt"
	"sort"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/rxdps93/home-recipes/internal/db"
)

func GenerateTagsHTML() string {
	head := GenerateHeadNode("Tags", "View All Tags")

	tags, err := db.GetAllTags()
	if err != nil {
		body := GenerateErrorNode(err, "Unable to load tags; please try again later.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}

	sections := make(map[rune][]elem.Node)
	for _, tag := range tags {
		sections[rune(tag[0])] = append(sections[rune(tag[0])],
			elem.Li(nil,
				elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/tags/%v", tag), attrs.Class: "link"},
					elem.Text(tag),
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

	body := GenerateBodyStructure("View All Recipe Tags",
		GenerateRecipeNavLinks(),
		GenerateJumpLinks(ltrs),
		GenerateJumpDestinations(ltrs, sections),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}
