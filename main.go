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
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	// router := gin.Default()
	// router.GET("/recipes", getRecipes)
	http.HandleFunc("/", Home)
	http.HandleFunc("/recipes", Recipes)

	fmt.Println("Connecting...")
	Connect()
	fmt.Println("Database opened")

	_, err := AddTestGrilledCheeseRecipe()
	if err != nil {
		log.Fatal(err)
	}

	_, err = AddTestChocolateMilkshakeRecipe()
	if err != nil {
		log.Fatal(err)
	}

	// rec, err := GetRecipeByID(recDB, ida)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// recs, err := GetAllRecipes(recDB)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, rec := range recs {
	// printRecipe(rec)
	// }

	// RemoveRecipe(recDB, testId)
	// router.Run("localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func cleanup() {
	log.Println("\nReceived Interrupt Event")
	log.Println("Test data cleanup (wiping database)")
	WipeDatabase()
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
}

// func getRecipes(c *gin.Context) {
// 	recs, err := GetAllRecipes(recDB)
// if err != nil {
// 	c.IndentedJSON(http.StatusInternalServerError, err)
// } else {
// 	c.IndentedJSON(http.StatusOK, recs)
// 	}
// }
