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

	router.Run("localhost:8080")

	fmt.Println("Connecting...")
	recDB := Connect()
	fmt.Println("Database opened")

	testId, err := AddTestRecipe(recDB)
	if err != nil {
		log.Fatal(err)
	}

	rec, err := QueryRecipeTableByID(recDB, testId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recipe found: %v\n", rec)

	DeleteTestRecipe(recDB, testId)
}

func getRecipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
