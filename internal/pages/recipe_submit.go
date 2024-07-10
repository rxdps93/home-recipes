package pages

import (
	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

// TODO: move instructions into separate func; look into expand/collapse of instr section
func GenerateRecipeSubmitHTML() string {
	head := GenerateHeadNode("Submit a Recipe", "Submit a Recipe", false)
	body := GenerateBodyStructure("Submit a Recipe",
		GenerateRecipeNavLinks(),
		elem.Div(attrs.Props{attrs.Class: "rs-instr"},
			elem.H2(nil, elem.Text("Instructions for Submission")),
			elem.P(nil,
				elem.Text("When submitting a recipe, the following fields are required:"),
				elem.Ol(nil,
					elem.Li(nil,
						elem.Text("Name"),
						elem.Ul(nil,
							elem.Li(nil, elem.Text("Should be brief and in natural language.")),
							elem.Li(nil, elem.Text("No two recipes should have the same name.")),
						),
					),
					elem.Li(nil,
						elem.Text("Description"),
						elem.Ul(nil,
							elem.Li(nil, elem.Text("Should give an overview of the recipe.")),
							elem.Li(nil, elem.Text("There is no character limit but should generally be a short paragraph at most.")),
						),
					),
					elem.Li(nil,
						elem.Text("Ingredients"),
						elem.Ul(nil,
							elem.Li(nil, elem.Text("Ingredients are unique in the database - this form will let you view existing ingredients.")),
						),
					),
					elem.Li(nil,
						elem.Text("Instructions"),
						elem.Ul(nil,
							elem.Li(nil, elem.Text("Each step can be as descriptive as you wish.")),
							elem.Li(nil, elem.Text("No need to number the step yourself - the site will do that automatically.")),
						),
					),
					elem.Li(nil,
						elem.Text("Tags"),
						elem.Ul(nil,
							elem.Li(nil, elem.Text("Much like ingredients duplicate tags should not be submitted.")),
							elem.Li(nil, elem.Text("There is no limit to how many tags can be tied to a single recipe.")),
						),
					),
					elem.Li(nil,
						elem.Text("Source"),
						elem.Ul(nil,
							elem.Li(nil, elem.Text("Do not enter personal information (e.g. names) here as the website will be public.")),
							elem.Li(nil, elem.Text("If the source is a website please enter the whole URL: (e.g. https://example.com/recipes/my-recipe.html)")),
							elem.Li(nil, elem.Text("If the source is a book do not include the author if the book is not publically available (e.g. a family recipe book)")),
						),
					),
				),
			),
		),
	)

	return elem.Html(nil, head, body).Render()
}
