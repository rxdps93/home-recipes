package pages

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GenerateHeadNode(title string, description string, useHTMX bool) elem.Node {
	content := elem.Head(nil,
		elem.Title(nil, elem.Text(title)),
		elem.Link(attrs.Props{attrs.Rel: "stylesheet", attrs.Type: "text/css", attrs.Href: "/assets/css/styles.css"}),
		elem.Meta(attrs.Props{attrs.Charset: "utf-8"}),
		elem.Meta(attrs.Props{attrs.Name: "description", attrs.Content: description}),
		elem.Meta(attrs.Props{attrs.Name: "viewport", attrs.Content: "width=device-width, initial-scale=1"}),
	)

	if useHTMX {
		content.Children = append(content.Children, elem.Script(attrs.Props{attrs.Src: "https://unpkg.com/htmx.org@2.0.0"}))
	}

	return content
}

func GenerateBodyStructure(headerText string, mainContent ...elem.Node) elem.Node {
	return elem.Body(nil,
		elem.Header(nil,
			elem.H1(nil, elem.Text(headerText))),
		generateNavigationHTML(),
		elem.Main(nil, mainContent...),
		// generateFooterHTML(),
	)
}

func generateNavigationHTML() elem.Node {
	return elem.Nav(nil,
		elem.Ul(nil,
			elem.Li(nil, elem.A(attrs.Props{attrs.Href: "/", attrs.Class: "link"}, elem.Text("Home"))),
			elem.Li(nil, elem.A(attrs.Props{attrs.Href: "/recipes", attrs.Class: "link"}, elem.Text("Recipes"))),
			elem.Li(nil, elem.A(attrs.Props{attrs.Href: "/test", attrs.Class: "link"}, elem.Text("Test Page"))),
		),
	)
}

// TODO: may not use this, just for experimenting
func generateFooterHTML() elem.Node {
	return elem.Footer(nil,
		elem.H3(nil, elem.Text("This is an example of a footer")),
	)
}

// TODO: revisit this in the future
func GenerateErrorNode(err error, msg string) elem.Node {
	return elem.Body(nil,
		generateNavigationHTML(),
		elem.H1(nil,
			elem.Text(msg),
		),
		elem.P(nil, elem.Text(err.Error())),
	)
}

func GenerateRecipeNavLinks() elem.Node {
	return elem.Div(attrs.Props{attrs.Class: "rec-nav"},
		elem.H3(nil, elem.Text("Recipe Page Navigation")),
		elem.A(attrs.Props{attrs.Href: "/recipes", attrs.Class: "link"},
			elem.Text("All Recipes"),
		),
		elem.A(attrs.Props{attrs.Href: "/tags", attrs.Class: "link"},
			elem.Text("Recipe Tags"),
		),
		elem.A(attrs.Props{attrs.Href: "/recipe-search", attrs.Class: "link"},
			elem.Text("Search Recipes"),
		),
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
