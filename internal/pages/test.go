package pages

import (
	"github.com/chasefleming/elem-go"
)

func GenerateTestHTML() string {
	head := GenerateHeadNode("CSS Test Page", "A Test Page For CSS")

	body := GenerateBodyStructure("CSS Test Page",
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
	)
	html := elem.Html(nil, head, body)
	return html.Render()
}
