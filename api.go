package main

import (
	"net/http"

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

func generateHomePage() string {
	head := elem.Head(nil,
		elem.Script(attrs.Props{attrs.Src: "https://unpkg.com/htmx.org@1.6.1"}),
	)

	body := elem.Div(nil,
		elem.H1(nil, elem.Text("This is the homepage!")),
		elem.A(attrs.Props{attrs.Href: "/recipes"}, elem.Text("View Recipes!")),
	)

	html := elem.Html(nil, head, body)

	return html.Render()
}

func generateRecipesPage() string {
	head := elem.Head(nil, elem.Title(nil, elem.Text("Recipes")))

	body := elem.Body(nil)

	html := elem.Html(nil, head, body)

	return html.Render()
}
