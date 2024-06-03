package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
)

type Ingredient struct {
	Label    string  `json:"label"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type Recipe struct {
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Instructions []string     `json:"instructions"`
	Ingredients  []Ingredient `json:"ingredients"`
	Tags         []string     `json:"tags"`
	Source       string       `json:"source"`
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", Home)
	mux.HandleFunc("GET /recipes", Recipes)
	mux.HandleFunc("GET /recipes/{id}", RecipeDetail)
	mux.HandleFunc("GET /tags", Tags)
	mux.HandleFunc("GET /tags/{tag}", RecipesByTag)
	mux.HandleFunc("GET /test", Test)

	log.Println("Connecting...")
	Connect()
	log.Println("Connected to Database")

	populateTestData()

	http.ListenAndServe(":8080", mux)
}

func populateTestData() {
	_, err := AddTestGrilledCheeseRecipe()
	if err != nil {
		log.Fatal(err)
	}

	_, err = AddTestChocolateMilkshakeRecipe()
	if err != nil {
		log.Fatal(err)
	}

	_, err = AddTestFreshGuacamoleRecipe()
	if err != nil {
		log.Fatal(err)
	}
}

func cleanup() {
	log.Println("\nReceived Interrupt Event")
	log.Println("Test data cleanup (wiping database)")
	WipeDatabase()
	Disconnect()
}

func printRecipe(rec Recipe) {
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
