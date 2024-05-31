package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

func Home(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := generateHomePage()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Recipes(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := generateRecipesPage()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RecipeDetail(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		id := req.PathValue("id")
		content := generateRecipeDetailPage(id)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func generateRecipeDetailPage(id string) string {
	recID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		// TODO: handle invalid id
		log.Fatal(err)
	}

	rec, err := GetRecipeByID(recID)
	if err != nil {
		// TODO: handle invalid id
		log.Fatal(err)
	}

	head := elem.Head(nil, elem.Title(nil, elem.Text(rec.Name)))

	ings := elem.Ul(nil)
	for _, ing := range rec.Ingredients {
		ings.Children = append(ings.Children,
			elem.Li(nil, elem.Text(fmt.Sprintf("%v %v %v", ing.Quantity, ing.Unit, ing.Label))),
		)
	}

	instr := elem.Ol(nil)
	for _, ins := range rec.Instructions {
		instr.Children = append(instr.Children,
			elem.Li(nil, elem.Text(ins)),
		)
	}

	body := elem.Body(nil,
		elem.Header(nil, elem.H1(nil, elem.Text(rec.Name))),
		elem.P(nil, elem.Text(rec.Description)),
		elem.Main(nil,
			elem.H2(nil, elem.Text("Ingredients")),
			ings,
			elem.H2(nil, elem.Text("Instructions")),
			instr,
		),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}

func generateHomePage() string {
	head := elem.Head(nil, elem.Title(nil, elem.Text("Home")))

	body := elem.Div(nil,
		elem.H1(nil, elem.Text("Test Homepage")),
		elem.A(attrs.Props{attrs.Href: "/recipes"}, elem.Text("View All Recipes")),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}

func generateRecipesPage() string {
	head := elem.Head(nil, elem.Title(nil, elem.Text("Recipes")))

	recs, err := GetAllRecipes()
	if err != nil {
		// TODO: handle this error case
		log.Fatal(err)
	}
	sort.Slice(recs, func(i, j int) bool {
		return recs[i].Name < recs[j].Name
	})

	body := elem.Body(nil)
	for _, rec := range recs {
		body.Children = append(body.Children,
			elem.H1(nil,
				elem.A(attrs.Props{attrs.Href: fmt.Sprintf("/recipes/%v", rec.ID)},
					elem.Text(rec.Name)),
			),
		)
	}

	html := elem.Html(nil, head, body)

	return html.Render()
}
