package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type RecipeDB struct {
	ID           int64  `json:"recipe_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
}

type IngredientDB struct {
	ID    int64  `json:"ingredient_id"`
	Label string `json:"label"`
}

type UnitDB struct {
	ID    int64  `json:"unit_id"`
	Label string `json:"label"`
}

type RecipeIngredientDB struct {
	ID           int64   `json:"rec_ing_id"`
	RecipeID     int64   `json:"recipe_id"`
	IngredientID int64   `json:"ingredient_id"`
	UnitID       int64   `json:"unit_id"`
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
	result, err := db.Exec("INSERT INTO recipe (name, description, instructions) VALUES (?, ?, ?)", "potato slop", "not very good", "eat|enjoy")
	if err != nil {
		return 0, fmt.Errorf("addTestRecipe: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addTestRecipe: %v", err)
	}
	return id, nil
}

func QueryRecipeTableByID(db *sql.DB, id int64) (RecipeDB, error) {
	var rec RecipeDB

	row := db.QueryRow("SELECT * FROM recipe WHERE recipe_id = ?", id)
	if err := row.Scan(&rec.ID, &rec.Name, &rec.Description, &rec.Instructions); err != nil {
		if err == sql.ErrNoRows {
			return rec, fmt.Errorf("queryRecipeTableByID %d: no such recipe", id)
		}
		return rec, fmt.Errorf("queryRecipeTableByID %d: %v", id, err)
	}
	return rec, nil
}

func QueryIngredientTableByID(db *sql.DB, id int64) (IngredientDB, error) {
	var ing IngredientDB

	row := db.QueryRow("SELECT * FROM ingredient WHERE ingredient_id = ?", id)
	if err := row.Scan(&ing.ID, &ing.Label); err != nil {
		if err == sql.ErrNoRows {
			return ing, fmt.Errorf("queryIngredientTableByID %d: no such ingredient", id)
		}
		return ing, fmt.Errorf("queryIngredientTableByID %d: %v", id, err)
	}
	return ing, nil
}

func QueryUnitTableByID(db *sql.DB, id int64) (UnitDB, error) {
	var unit UnitDB

	row := db.QueryRow("SELECT * FROM unit WHERE unit_id = ?", id)
	if err := row.Scan(&unit.ID, &unit.Label); err != nil {
		if err == sql.ErrNoRows {
			return unit, fmt.Errorf("queryUnitTableByID %d: no such unit", id)
		}
		return unit, fmt.Errorf("queryUnitTableByID %d: %v", id, err)
	}
	return unit, nil
}

func QueryRecipeIngredientsByID(db *sql.DB, recID int64) ([]RecipeIngredientDB, error) {
	var ings []RecipeIngredientDB

	rows, err := db.Query("SELECT * FROM recipe_ingredient WHERE recipe_id = ?", recID)
	if err != nil {
		return nil, fmt.Errorf("queryRecipeIngredientsByID %q: %v", recID, err)
	}
	defer rows.Close()
	for rows.Next() {
		var ri RecipeIngredientDB
		if err := rows.Scan(&ri.ID, &ri.RecipeID, &ri.IngredientID, &ri.UnitID, &ri.Quantity); err != nil {
			return nil, fmt.Errorf("queryRecipeIngredientsByID %q: %v", recID, err)
		}
		ings = append(ings, ri)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryRecipeIngredientsByID %q: %v", recID, err)
	}
	return ings, nil
}

func GetRecipeByID(db *sql.DB, id int64) (RecipeFull, error) {
	var rec RecipeFull

	recDB, err := QueryRecipeTableByID(db, id)
	if err != nil {
		return rec, err
	}

	riDB, err := QueryRecipeIngredientsByID(db, id)
	if err != nil {
		return rec, err
	}

	rec.id = id
	rec.name = recDB.Name
	rec.desc = recDB.Description
	rec.inst = strings.Split(recDB.Instructions, "|")
	rec.ingr = make(map[string]string)

	for _, ri := range riDB {
		unitDB, err := QueryUnitTableByID(db, ri.UnitID)
		if err != nil {
			return rec, err
		}

		ingDB, err := QueryIngredientTableByID(db, ri.IngredientID)
		if err != nil {
			return rec, err
		}

		rec.ingr[ingDB.Label] = fmt.Sprintf("%v %v", ri.Quantity, unitDB.Label)
	}

	return rec, nil
}
