package pages

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GenerateTestHTML() string {
	head := elem.Head(nil,
		elem.Title(nil, elem.Text("css test page")),
		elem.Link(attrs.Props{attrs.Rel: "stylesheet", attrs.Type: "text/css", attrs.Href: "/css/style.css"}),
		elem.Meta(attrs.Props{attrs.Charset: "utf-8"}),
		elem.Meta(attrs.Props{attrs.Name: "description", attrs.Content: "a test page for css"}),
		elem.Meta(attrs.Props{attrs.Name: "viewport", attrs.Content: "width=device-width, initial-scale=1"}),
	)

	body := elem.Body(nil,
		elem.Header(nil, elem.H1(nil, elem.Text("CSS Test Page"))),
		elem.Nav(nil, elem.Ul(nil,
			elem.A(attrs.Props{attrs.Href: "/"}, elem.Li(nil, elem.Text("Home"))),
			elem.A(attrs.Props{attrs.Href: "/recipes"}, elem.Li(nil, elem.Text("Recipes"))),
			elem.A(attrs.Props{attrs.Href: "/tags"}, elem.Li(nil, elem.Text("Tags"))),
		)),
		elem.Main(nil,
			elem.H1(nil, elem.Text("What does an H1 look like here?")),
			elem.H2(nil, elem.Text("This is for testing stylesheet linking :^)")),
			elem.H2(nil, elem.Text("I hope this is much easier than stylemanager!")),
			elem.P(nil, elem.Text("If so this will make development much less tedious.")),
			elem.H2(nil, elem.Text("List for testing:")),
			elem.Ul(nil,
				elem.Li(nil, elem.Text("List item 1")),
				elem.Li(nil, elem.Text("Nice weather today")),
				elem.Li(nil, elem.Text("A Møøse once bit my sister...")),
			),
		),
	)
	html := elem.Html(nil, head, body)
	return html.Render()
}
