package api

import (
	"fmt"
	"net/http"

	"github.com/chasefleming/elem-go"
	"github.com/rxdps93/home-recipes/internal/pages"
)

// TODO: better handling of 'method not allowed' scenarios

func HomePage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		if req.URL.Path != "/" {
			content := pages.GenerateNotFoundHTML()
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(content))
		} else {
			content := pages.GenerateHomeHTML()
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(content))
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RecipesPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := pages.GenerateRecipesHTML()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RecipeDetailPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		id := req.PathValue("id")
		content := pages.GenerateRecipeDetailHTML(id)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RecipeSearchPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := pages.GenerateRecipeSearchHTML()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Search(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.ParseForm()

		nq := req.FormValue("name")
		tq := req.Form["tags"]

		content := pages.GenerateTableBody(nq, tq)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// TODO: this
func Submit(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.ParseForm()

		name := req.FormValue("name")
		desc := req.FormValue("description")

		content := elem.Div(nil,
			elem.P(nil, elem.Text(fmt.Sprintf("Name: %v", name))),
			elem.P(nil, elem.Text(fmt.Sprintf("Description: %v", desc))),
		).Render()

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK) // TODO: set this based on submission results
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func TagsPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := pages.GenerateTagsHTML()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RecipesByTagPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		tag := req.PathValue("tag")
		content := pages.GenerateRecipesByTagHTML(tag)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RecipeSubmitPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := pages.GenerateRecipeSubmitHTML()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	}
}

func TestPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := pages.GenerateTestHTML()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SearchTest(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.ParseForm()

		nq := req.FormValue("name")
		tq := req.Form["tags"]

		content := pages.GenerateTestTableBody(nq, tq)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
