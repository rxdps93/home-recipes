package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type RecipeFull struct {
	name string
	desc string
	inst map[uint]string
	ingr map[string]string
}

func main() {
	fmt.Println("Connecting...")
	recDB := Connect()
	fmt.Println("Database opened")

	testId, err := AddTestRecipe(recDB)
	if err != nil {
		log.Fatal(err)
	}

	rec, err := QueryRecipeTable(recDB, testId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recipe found: %v\n", rec)

	DeleteTestRecipe(recDB, testId)
}
