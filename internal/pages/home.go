package pages

import "github.com/chasefleming/elem-go"

func GenerateHomeHTML() string {
	head := GenerateHeadNode("Home", "Home Description", false)

	body := GenerateBodyStructure("Test Homepage",
		elem.H2(nil, elem.Text("This website is under construction :^)")),
		elem.H2(nil, elem.Text("Styling is via traditional css stylesheet linking.")),
		elem.P(nil, elem.Text("Stylemanager was very interesting to learn but made complex styling more tedious than just writing CSS.")),
		elem.H2(nil, elem.Text("Here is a list:")),
		elem.Ul(nil,
			elem.Li(nil, elem.Text("List item 1")),
			elem.Li(nil, elem.Text("Nice weather today")),
			elem.Li(nil, elem.Text("A Møøse once bit my sister...")),
		),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}
