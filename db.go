package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type RecipeDB struct {
	ID           int    `json:"recipe_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
}

type IngredientDB struct {
	ID    int    `json:"ingredient_id"`
	Label string `json:"label"`
}

type UnitDB struct {
	ID    int    `json:"unit_id"`
	Label string `json:"label"`
}

type RecipeIngredientDB struct {
	ID           int     `json:"rec_ing_id"`
	RecipeID     int     `json:"recipe_id"`
	IngredientID int     `json:"ingredient_id"`
	UnitID       int     `json:"unit_id"`
	Quantity     float64 `json:"quantity"`
}

func Connect() *sql.DB {
	DB, err := sql.Open("sqlite3", "recipes.db")
	if err != nil {
		log.Fatal(err)
	}

	return DB
}

func DeleteTestRecipe(db *sql.DB, id int64) {
	_, err := db.Exec("DELETE FROM recipe WHERE recipe_id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

func AddTestRecipe(db *sql.DB) (int64, error) {
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

func QueryRecipeTable(db *sql.DB, id int64) (RecipeDB, error) {
	var rec RecipeDB

	row := db.QueryRow("SELECT * FROM recipe WHERE recipe_id = ?", id)
	if err := row.Scan(&rec.ID, &rec.Name, &rec.Description, &rec.Instructions); err != nil {
		if err == sql.ErrNoRows {
			return rec, fmt.Errorf("queryRecipeTable %d: no such recipe", id)
		}
		return rec, fmt.Errorf("queryRecipeTable %d: %v", id, err)
	}
	return rec, nil
}
