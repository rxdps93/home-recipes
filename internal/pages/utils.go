package pages

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GenerateHeadNode(title string, description string) elem.Node {
	return elem.Head(nil,
		elem.Title(nil, elem.Text(title)),
		elem.Link(attrs.Props{attrs.Rel: "stylesheet", attrs.Type: "text/css", attrs.Href: "/css/style.css"}),
		elem.Meta(attrs.Props{attrs.Charset: "utf-8"}),
		elem.Meta(attrs.Props{attrs.Name: "description", attrs.Content: description}),
		elem.Meta(attrs.Props{attrs.Name: "viewport", attrs.Content: "width=device-width, initial-scale=1"}),
	)
}

func GenerateBodyStructure(headerText string, mainContent ...elem.Node) elem.Node {
	return elem.Body(nil,
		elem.Header(nil,
			elem.H1(nil, elem.Text(headerText))),
		GenerateNavigationHTML(),
		elem.Main(nil, mainContent...),
	)
}

// TODO: consider moving tags to be within the recipes section
// TODO: swap A and Li to be valid html
func GenerateNavigationHTML() elem.Node {
	return elem.Nav(nil,
		elem.Ul(nil,
			elem.A(attrs.Props{attrs.Href: "/"}, elem.Li(nil, elem.Text("Home"))),
			elem.A(attrs.Props{attrs.Href: "/recipes"}, elem.Li(nil, elem.Text("Recipes"))),
			elem.A(attrs.Props{attrs.Href: "/tags"}, elem.Li(nil, elem.Text("Tags"))),
		),
	)
}

func GenerateErrorNode(err error, msg string) elem.Node {
	return elem.Body(nil,
		GenerateNavigationHTML(),
		elem.H1(nil,
			elem.Text(msg),
		),
		elem.P(nil, elem.Text(err.Error())),
	)
}

func GenerateJumpLinks(ltrs []rune) elem.Node {
	content := elem.Div(attrs.Props{attrs.Class: "jump-link"},
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

func GenerateJumpDestinations(ltrs []rune, sections map[rune][]elem.Node) elem.Node {
	content := elem.Div(attrs.Props{attrs.Class: "jump-dest"})

	for _, ltr := range ltrs {
		content.Children = append(content.Children,
			elem.Div(attrs.Props{attrs.Class: "jump-sec", attrs.ID: string(ltr)},
				elem.H3(nil, elem.Text(string(ltr))),
				elem.Ul(nil, sections[ltr]...),
			),
		)
	}

	return content
}
