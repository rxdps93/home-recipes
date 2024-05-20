package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type recipeDB struct {
	ID           int    `json:"recipe_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
}

type ingredientDB struct {
	ID    int    `json:"ingredient_id"`
	Label string `json:"label"`
}

type unitDB struct {
	ID    int    `json:"unit_id"`
	Label string `json:"label"`
}

type recipeIngredientDB struct {
	ID           int     `json:"rec_ing_id"`
	RecipeID     int     `json:"recipe_id"`
	IngredientID int     `json:"ingredient_id"`
	UnitID       int     `json:"unit_id"`
	Quantity     float64 `json:"quantity"`
}

type recipeFull struct {
	name string
	desc string
	inst map[uint]string
	ingr map[string]string
}

var db *sql.DB

func main() {
	var err error
	fmt.Println("Opening database...")
	db, err = sql.Open("sqlite3", "recipes.db")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database opened")

	testId, err := addTestRecipe()
	if err != nil {
		log.Fatal(err)
	}

	rec, err := queryRecipeTable(testId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recipe found: %v\n", rec)

	deleteTestRecipe(testId)
}

func deleteTestRecipe(id int64) {
	_, err := db.Exec("DELETE FROM recipe WHERE recipe_id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

func addTestRecipe() (int64, error) {
	result, err := db.Exec("INSERT INTO recipe (name, description, instructions) VALUES (?, ?, ?)", "potato slop", "not very good", "1) eat\n2)poop")
	if err != nil {
		return 0, fmt.Errorf("addTestRecipe: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addTestRecipe: %v", err)
	}
	return id, nil
}

func queryRecipeTable(id int64) (recipeDB, error) {
	var rec recipeDB

	row := db.QueryRow("SELECT * FROM recipe WHERE recipe_id = ?", id)
	if err := row.Scan(&rec.ID, &rec.Name, &rec.Description, &rec.Instructions); err != nil {
		if err == sql.ErrNoRows {
			return rec, fmt.Errorf("queryRecipeTable %d: no such recipe", id)
		}
		return rec, fmt.Errorf("queryRecipeTable %d: %v", id, err)
	}
	return rec, nil
}
