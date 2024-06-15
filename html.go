package main

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func generateErrorNode(err error, msg string) elem.Node {
	return elem.Body(attrs.Props{attrs.Style: BodyStyle.ToInline()},
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
		elem.Meta(attrs.Props{attrs.Charset: "utf-8"}),
		elem.Meta(attrs.Props{attrs.Name: "description", attrs.Content: description}),
		elem.Meta(attrs.Props{attrs.Name: "viewport", attrs.Content: "width=device-width, initial-scale=1"}),
	)
}

func generateBodyStructure(headerText string, mainContent ...elem.Node) elem.Node {
	return elem.Body(
		attrs.Props{attrs.Style: BodyStyle.ToInline()},
		elem.Header(nil,
			elem.H1(attrs.Props{attrs.Style: HeaderH1Style.ToInline()},
				elem.Text(headerText))),
		generateNavigationHTML(),
		elem.Main(attrs.Props{attrs.Style: MainStyle.ToInline()}, mainContent...),
	)
}

func GenerateHomeHTML() string {
	head := generateHeadNode("Home", "Home Description")

	body := generateBodyStructure("Test Homepage",
		elem.H2(attrs.Props{attrs.Style: BaseH2Style.ToInline()},
			elem.Text("Currently Under Construction :^)")),
		elem.P(nil, elem.Text("I can't wait to see how this develops!")),
	)
	// body := elem.Div(
	// 	attrs.Props{attrs.Style: BodyStyle.ToInline()},
	// generateNavigationHTML(),
	// elem.H1(attrs.Props{attrs.Style: BaseH1Style.ToInline()}, elem.Text("Test Homepage")),
	// 	elem.A(attrs.Props{attrs.Href: "/recipes"}, elem.Text("View All Recipes")),
	// )

	html := elem.Html(nil, head, body)

	return html.RenderWithOptions(elem.RenderOptions{StyleManager: StyleMgr})
}

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

	body := elem.Body(attrs.Props{attrs.Style: BodyStyle.ToInline()},
		generateNavigationHTML(),
		elem.H2(nil, elem.Text("All Recipes:")),
		elem.Ul(nil,
			elem.TransformEach(recs, func(rec Recipe) elem.Node {
				return elem.Li(nil,
					elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/recipes/%v", rec.ID)},
						elem.Text(rec.Name),
					),
				)
			})...,
		),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}

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

	body := elem.Body(attrs.Props{attrs.Style: BodyStyle.ToInline()},
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

	body := elem.Body(attrs.Props{attrs.Style: BodyStyle.ToInline()},
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

	body := elem.Body(attrs.Props{attrs.Style: BodyStyle.ToInline()},
		generateNavigationHTML(),
		elem.H2(nil, elem.Text(fmt.Sprintf("Recipes Tagged As %v:", tag))),
		tags,
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}

func generateNavigationHTML() elem.Node {
	navLiClass := StyleMgr.AddCompositeStyle(NavLiMediaQuery)
	navAHoverClass := StyleMgr.AddCompositeStyle(NavAHoverLiClass)
	navMediaQuery := StyleMgr.AddCompositeStyle(NavMediaQuery)
	navBeforeMediaQuery := StyleMgr.AddCompositeStyle(NavBeforeMediaQuery)
	return elem.Nav(attrs.Props{
		attrs.Style: NavStyle.ToInline(),
		attrs.Class: navMediaQuery + " " + navBeforeMediaQuery,
	},
		elem.Ul(attrs.Props{attrs.Style: NavUlStyle.ToInline()},
			elem.Li(attrs.Props{
				attrs.Style: NavLiStyle.ToInline(),
				attrs.Class: navLiClass,
			},
				elem.A(attrs.Props{
					attrs.Href:  "/",
					attrs.Class: navAHoverClass,
				}, elem.Text("Home"))),

			elem.Li(attrs.Props{
				attrs.Style: NavLiStyle.ToInline(),
				attrs.Class: navLiClass,
			},
				elem.A(attrs.Props{
					attrs.Href:  "/recipes",
					attrs.Class: navAHoverClass,
				}, elem.Text("Recipes"))),

			elem.Li(attrs.Props{
				attrs.Style: NavLiStyle.ToInline(),
				attrs.Class: navLiClass,
			},
				elem.A(attrs.Props{
					attrs.Href:  "/tags",
					attrs.Class: navAHoverClass,
				}, elem.Text("Tags"))),
		),
	)
}
