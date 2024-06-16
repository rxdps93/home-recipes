package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := GenerateHomeHTML()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Recipes(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := GenerateRecipesHTML()
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
		content := GenerateRecipeDetailHTML(id)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Tags(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := GenerateTagsHTML()
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
		content := GenerateRecipesByTagHTML(tag)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Test(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		content := GenerateTestHTML()
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
