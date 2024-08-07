package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rxdps93/home-recipes/internal/api"
	"github.com/rxdps93/home-recipes/internal/db"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", api.HomePage)
	mux.HandleFunc("GET /recipes", api.RecipesPage)
	mux.HandleFunc("GET /recipes/{id}", api.RecipeDetailPage)
	mux.HandleFunc("GET /tags", api.TagsPage)
	mux.HandleFunc("GET /tags/{tag}", api.RecipesByTagPage)
	mux.HandleFunc("GET /recipe-search", api.RecipeSearchPage)
	mux.HandleFunc("GET /recipe-submit", api.RecipeSubmitPage)

	mux.HandleFunc("POST /search", api.Search)
	mux.HandleFunc("POST /submit", api.Submit)

	mux.HandleFunc("GET /test", api.TestPage)
	mux.HandleFunc("POST /search-test", api.SearchTest)

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Println("Connecting...")
	db.Connect()
	log.Println("Connected to Database")

	populateTestData()

	http.ListenAndServe(":8080", mux)
}

func populateTestData() {
	_, err := db.AddTestGrilledCheeseRecipe()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.AddTestChocolateMilkshakeRecipe()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.AddTestFreshGuacamoleRecipe()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.AddTestGarlicLimeChicken()
	if err != nil {
		log.Fatal(err)
	}
}

func cleanup() {
	log.Println("\nReceived Interrupt Event")
	log.Println("Test data cleanup (wiping database)")
	db.WipeDatabase()
	db.Disconnect()
}

func printRecipe(rec db.Recipe) {
	fmt.Printf("Recipe Name:\n\t%v\n", rec.Name)
	fmt.Printf("Description:\n\t%v\n", rec.Description)
	fmt.Printf("Ingredients:\n")
	for _, ing := range rec.Ingredients {
		fmt.Printf("\t%v:\t%v %v\n", ing.Label, ing.Quantity, ing.Unit)
	}
	fmt.Printf("Instructions:\n")
	for i, step := range rec.Instructions {
		fmt.Printf("\t%v) %v\n", i, step)
	}
	fmt.Printf("Tags:\n")
	for _, tag := range rec.Tags {
		fmt.Printf("\t%v\n", tag)
	}
	fmt.Printf("Recipe Source:\n\t%v\n", rec.Source)
}
