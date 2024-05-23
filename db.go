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
	rec := RecipeFull{
		name: "Grilled Cheese Sandwich",
		desc: "A classic and simple sandwich",
		inst: []string{
			"Spread butter onto one side of each slice of bread",
			"Put a skillet on the stove on medium/low heat.",
			"Place one slice of bread into skillet, butter side down.",
			"Place cheese on top of the bread in the skillet.",
			"Place the remaining slice of bread on the cheese, butter side up.",
			"Cover skillet and wait until bottom slice is golden brown.",
			"Carefully flip and cover.",
			"Once golden brown and cheese is adequately melted, the sandwich is ready.",
		},
		ingr: map[string]string{
			"Bread":  "2 Slices",
			"Cheese": "4 Slices",
			"Butter": "0 x",
		},
	}

	id, err := SubmitRecipe(db, rec)
	if err != nil {
		return 0, err
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

func SubmitRecipe(db *sql.DB, rec RecipeFull) (int64, error) {
	recStmt, err := db.Prepare(
		"INSERT INTO recipe (name, description, instructions) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	ingStmt, err := db.Prepare("INSERT INTO ingredient (label) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}

	unitStmt, err := db.Prepare("INSERT INTO unit (label) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}

	riStmt, err := db.Prepare("INSERT INTO recipe_ingredient (recipe_id, ingredient_id, unit_id, quantity) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	/*
	 * For Recipe:
	 * > name            -> rec.name
	 * > description     -> rec.disc
	 * > instructions    -> strings.join(rec.inst, "|")
	 * exec recstmt with above, store rid
	 * For Ingredient and Unit:
	 * for each k,v in rec.ingr
	 * ilabel = k
	 * ulabel = strings.split(v, " ")[1]
	 * quantity = strings.split(v, " ")[0]
	 * exec ingstmt with ilabel, store iid
	 * exec unitstmt with ulabel, store uid
	 * exec ristmt with rid, iid, uid, quantity
	 */
	// RECIPE SECTION
	result, err := recStmt.Exec(rec.name, rec.desc, strings.Join(rec.inst, "|"))
	if err != nil {
		return 0, fmt.Errorf("SubmitRecipe: %v", err)
	}
	recID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("SubmitRecipe: %v", err)
	}

	// ING / UNIT SECTION
	for k, v := range rec.ingr {
		u := strings.Split(v, " ")

		// INGREDIENT
		result, err = ingStmt.Exec(k)
		if err != nil {
			return 0, fmt.Errorf("SubmitRecipe: %v", err)
		}
		ingID, err := result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("SubmitRecipe: %v", err)
		}

		// UNIT
		result, err = unitStmt.Exec(u[1])
		if err != nil {
			return 0, fmt.Errorf("SubmitRecipe: %v", err)
		}
		unitID, err := result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("SubmitRecipe: %v", err)
		}

		// RECIPE_INGREDIENT
		result, err = riStmt.Exec(recID, ingID, unitID, u[0])
		if err != nil {
			return 0, fmt.Errorf("SubmitRecipe: %v", err)
		}
		_, err = result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("SubmitRecipe: %v", err)
		}
	}

	return recID, nil
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
