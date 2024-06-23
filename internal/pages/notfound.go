package pages

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func GenerateNotFoundHTML() string {
	head := GenerateHeadNode("404 Page Not Found", "404 Page Not Found")

	body := GenerateBodyStructure("404 - Page Not Found",
		elem.Div(attrs.Props{attrs.Style: "width: 100%"},
			elem.Img(attrs.Props{
				attrs.Src:   "/assets/404.gif",
				attrs.Alt:   "Ah, ah, ah!",
				attrs.Style: "float: left; margin: auto; display: block;",
			}),
			elem.H3(attrs.Props{attrs.Style: "color: red"},
				elem.Text("####### Ah, ah, ah! That page wasn't found! #######"),
			),
			elem.A(attrs.Props{attrs.Href: "/", attrs.Class: "link"},
				elem.Text("Return Home"),
			),
		),
	)

	return elem.Html(nil, head, body).Render()
}
