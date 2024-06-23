package api

import (
	"net/http"

	"github.com/rxdps93/home-recipes/internal/pages"
)

func Home(w http.ResponseWriter, req *http.Request) {
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

func Recipes(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := pages.GenerateRecipesHTML()
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
		content := pages.GenerateRecipeDetailHTML(id)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Tags(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := pages.GenerateTagsHTML()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func RecipesByTag(w http.ResponseWriter, req *http.Request) {
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

func Test(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := pages.GenerateTestHTML()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
