package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type RecipeFull struct {
	id   int64
	name string
	desc string
	inst []string
	ingr map[string]string
}

func main() {
	router := gin.Default()
	router.GET("/recipes", getRecipes)

	// router.Run("localhost:8080")

	fmt.Println("Connecting...")
	recDB := Connect()
	fmt.Println("Database opened")

	testId, err := AddTestRecipe(recDB)
	if err != nil {
		log.Fatal(err)
	}

	rec, err := GetRecipeByID(recDB, testId)
	if err != nil {
		log.Fatal(err)
	}

	printRecipe(rec)

	RemoveRecipe(recDB, testId)
}

func printRecipe(rec RecipeFull) {
	fmt.Printf("Recipe Name:\n\t%v\n", rec.name)
	fmt.Printf("Description:\n\t%v\n", rec.desc)
	fmt.Printf("Ingredients:\n")
	for ing, amt := range rec.ingr {
		fmt.Printf("\t%v:\t%v\n", ing, amt)
	}
	fmt.Printf("Instructions:\n")
	for i, step := range rec.inst {
		fmt.Printf("\t%v) %v\n", i, step)
	}
}

func getRecipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
