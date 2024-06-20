package pages

import (
	"fmt"
	"sort"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/rxdps93/home-recipes/internal/db"
)

// TODO: update html structure & styling
func GenerateTagsHTML() string {
	head := GenerateHeadNode("Tags", "View All Tags")

	tags, err := db.GetAllTags()
	if err != nil {
		body := GenerateErrorNode(err, "Unable to load tags; please try again later.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}
	sort.Slice(tags, func(i, j int) bool {
		return tags[i] < tags[j]
	})

	body := elem.Body(nil,
		GenerateNavigationHTML(),
		elem.H2(nil, elem.Text("All Tags:")),
		elem.Ul(nil,
			elem.TransformEach(tags, func(tag string) elem.Node {
				return elem.Li(nil,
					elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/tags/%v", tag)},
						elem.Text(tag),
					),
				)
			})...,
		),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}
