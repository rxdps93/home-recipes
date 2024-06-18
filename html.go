package main

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"

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

func generateErrorNode(err error, msg string) elem.Node {
	return elem.Body(nil,
		generateNavigationHTML(),
		elem.H1(nil,
			elem.Text(msg),
		),
		elem.P(nil, elem.Text(err.Error())),
	)
}

func generateHeadNode(title string, description string) elem.Node {
	return elem.Head(nil,
		elem.Title(nil, elem.Text(title)),
		elem.Link(attrs.Props{attrs.Rel: "stylesheet", attrs.Type: "text/css", attrs.Href: "/css/style.css"}),
		elem.Meta(attrs.Props{attrs.Charset: "utf-8"}),
		elem.Meta(attrs.Props{attrs.Name: "description", attrs.Content: description}),
		elem.Meta(attrs.Props{attrs.Name: "viewport", attrs.Content: "width=device-width, initial-scale=1"}),
	)
}

func generateBodyStructure(headerText string, mainContent ...elem.Node) elem.Node {
	return elem.Body(nil,
		elem.Header(nil,
			elem.H1(nil, elem.Text(headerText))),
		generateNavigationHTML(),
		elem.Main(nil, mainContent...),
	)
}

func GenerateHomeHTML() string {
	head := generateHeadNode("Home", "Home Description")

	body := generateBodyStructure("Test Homepage",
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

func recipeJumpLinks() elem.Node {
	content := elem.Div(attrs.Props{attrs.Class: "rec-jump"},
		elem.Hr(nil),
		elem.H3(nil, elem.Text("Jump to Section...")),
	)

	for ltr := 'A'; ltr < 'Z'; ltr++ {
		content.Children = append(content.Children,
			elem.A(attrs.Props{attrs.Href: "#" + string(ltr)}, elem.Text(string(ltr))),
			elem.Text("&middot;"),
		)
	}
	content.Children = append(content.Children,
		elem.A(attrs.Props{attrs.Href: "#Z"}, elem.Text("Z")),
		elem.Hr(nil),
	)

	return content
}

// TODO: update styling on li
func GenerateRecipesHTML() string {
	head := generateHeadNode("Recipes", "View All Recipes")

	recs, err := GetAllRecipes()
	if err != nil {
		body := generateErrorNode(err, "Unable to load recipes; please try again later.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}
	sort.Slice(recs, func(i, j int) bool {
		return recs[i].Name < recs[j].Name
	})

	body := generateBodyStructure("Family Recipe Index",
		elem.H1(nil, elem.Text("View All Recipes")),
		recipeJumpLinks(),
		elem.Div(attrs.Props{attrs.Class: "rec-links"},
			elem.Ul(nil,
				elem.TransformEach(recs, func(rec Recipe) elem.Node {
					return elem.Li(nil,
						elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/recipes/%v", rec.ID)},
							elem.Text(rec.Name),
						),
					)
				})...,
			),
		),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}

// TODO: same as above
func GenerateTagsHTML() string {
	head := generateHeadNode("Tags", "View All Tags")

	tags, err := GetAllTags()
	if err != nil {
		body := generateErrorNode(err, "Unable to load tags; please try again later.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}
	sort.Slice(tags, func(i, j int) bool {
		return tags[i] < tags[j]
	})

	body := elem.Body(nil,
		generateNavigationHTML(),
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

// TODO: same as above
func GenerateRecipeDetailHTML(id string) string {
	recID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		head := generateHeadNode("Error", "Unable to Load Recipe")
		body := generateErrorNode(err, "Malformed recipe ID.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}

	rec, err := GetRecipeByID(recID)
	if err != nil {
		head := generateHeadNode("Error", "Unable to Load Recipe")
		body := generateErrorNode(err, "Unable to load recipe.")
		html := elem.Html(nil, head, body)
		return html.Render()
	}

	head := generateHeadNode(rec.Name, rec.Description)

	ings := elem.Ul(nil,
		elem.TransformEach(rec.Ingredients, func(ing Ingredient) elem.Node {
			return elem.Li(nil, elem.Text(fmt.Sprintf("%v %v %v", ing.Quantity, ing.Unit, ing.Label)))
		})...,
	)

	instr := elem.Ol(nil,
		elem.TransformEach(rec.Instructions, func(step string) elem.Node {
			return elem.Li(nil, elem.Text(step))
		})...,
	)

	tags := elem.Ul(nil,
		elem.TransformEach(rec.Tags, func(tag string) elem.Node {
			return elem.Li(nil,
				elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/tags/%v", tag)},
					elem.Text(tag)),
			)
		})...,
	)

	var src elem.Node
	_, err = url.ParseRequestURI(rec.Source)
	if err != nil {
		src = elem.Text(fmt.Sprintf(rec.Source))
	} else {
		src = elem.A(attrs.Props{attrs.Href: rec.Source}, elem.Text(rec.Source))
	}

	body := elem.Body(nil,
		generateNavigationHTML(),
		elem.Header(nil, elem.H1(nil, elem.Text(rec.Name))),
		elem.P(nil, elem.Text(rec.Description)),
		elem.Main(nil,
			elem.H2(nil, elem.Text("Ingredients")),
			ings,
			elem.H2(nil, elem.Text("Instructions")),
			instr,
			elem.H3(nil, elem.Text("Tags")),
			tags,
			elem.B(nil, elem.Text("Recipe Source: ")),
			src,
		),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}

// TODO: same as above
func GenerateRecipesByTagHTML(tag string) string {
	recs, err := GetAllRecipesForTagName(tag)
	if err != nil {
		head := generateHeadNode("Error", "Unable to load recipes")
		body := generateErrorNode(err, "Unable to load recipes")
		html := elem.Html(nil, head, body)
		return html.Render()
	}

	head := generateHeadNode(fmt.Sprintf("Recipes By Tag: %v", tag), fmt.Sprintf("Recipes Tagged As %v", tag))

	sort.Slice(recs, func(i, j int) bool {
		return recs[i].Name < recs[j].Name
	})

	tags := elem.Ul(nil,
		elem.TransformEach(recs, func(rec Recipe) elem.Node {
			return elem.Li(nil,
				elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/recipes/%v", rec.ID)},
					elem.Text(rec.Name),
				),
			)
		})...,
	)

	body := elem.Body(nil,
		generateNavigationHTML(),
		elem.H2(nil, elem.Text(fmt.Sprintf("Recipes Tagged As %v:", tag))),
		tags,
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}

// TODO: consider moving tags to be within the recipes section
// TODO: swap A and Li to be valid html
func generateNavigationHTML() elem.Node {
	return elem.Nav(nil,
		elem.Ul(nil,
			elem.A(attrs.Props{attrs.Href: "/"}, elem.Li(nil, elem.Text("Home"))),
			elem.A(attrs.Props{attrs.Href: "/recipes"}, elem.Li(nil, elem.Text("Recipes"))),
			elem.A(attrs.Props{attrs.Href: "/tags"}, elem.Li(nil, elem.Text("Tags"))),
		),
	)
}
