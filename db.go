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

func RemoveRecipe(db *sql.DB, id int64) {
	_, err := db.Exec("DELETE FROM recipe WHERE recipe_id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM recipe_ingredient WHERE recipe_id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

// FOR TESTING/DEVELOPMENT PURPOSES ONLY
// TODO: REMOVE WHEN NO LONGER NEEDED
func WipeDatabase(db *sql.DB) {
	db.Exec("DELETE FROM recipe")
	db.Exec("DELETE FROM ingredient")
	db.Exec("DELETE FROM unit")
	db.Exec("DELETE FROM recipe_ingredient")
}

func AddTestChocolateMilkshakeRecipe(db *sql.DB) (int64, error) {
	rec := Recipe{
		Name:        "Chocolate Milkshake",
		Description: "A classic favorite cold dessert.",
		Instructions: []string{
			"Add ice cream to blender.",
			"Pour milk into blender.",
			"Add ovaltine to blender.",
			"Blend until fully mixed.",
			"To get desired consistency add ice cream to thicken or milk to thin.",
			"To get desired flavor add ovaltine to taste.",
		},
		Ingredients: []Ingredient{
			{
				Label:    "Ice Cream",
				Quantity: 4,
				Unit:     "Scoops",
			},
			{
				Label:    "Milk",
				Quantity: 0.5,
				Unit:     "Cups",
			},
			{
				Label:    "Chocolate Malt Ovaltine",
				Quantity: 4,
				Unit:     "Tablespoons",
			},
		},
	}

	id, err := SubmitRecipe(db, rec)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func AddTestGrilledCheeseRecipe(db *sql.DB) (int64, error) {
	rec := Recipe{
		Name:        "Grilled Cheese Sandwich",
		Description: "A classic and simple sandwich.",
		Instructions: []string{
			"Spread butter onto one side of each slice of bread.",
			"Put a skillet on the stove on medium/low heat.",
			"Place one slice of bread into skillet, butter side down.",
			"Place cheese on top of the bread in the skillet.",
			"Place the remaining slice of bread on the cheese, butter side up.",
			"Cover skillet and wait until bottom slice is golden brown.",
			"Carefully flip and cover.",
			"Once golden brown and cheese is adequately melted, the sandwich is ready.",
		},
		Ingredients: []Ingredient{
			{
				Label:    "Bread",
				Quantity: 2,
				Unit:     "Slices",
			},
			{
				Label:    "Cheese",
				Quantity: 4,
				Unit:     "Slices",
			},
			{
				Label:    "Butter",
				Quantity: 1,
				Unit:     "Pat",
			},
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

func QueryIngredientTableByName(db *sql.DB, label string) (IngredientDB, error) {
	var ing IngredientDB

	row := db.QueryRow("SELECT * FROM ingredient WHERE label = ?", label)
	if err := row.Scan(&ing.ID, &ing.Label); err != nil {
		if err == sql.ErrNoRows {
			return ing, fmt.Errorf("QueryIngredientTableByName %v: no such ingredient", label)
		}
		return ing, fmt.Errorf("QueryIngredientTableByName %v: %v", label, err)
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

func QueryUnitTableByName(db *sql.DB, label string) (UnitDB, error) {
	var unit UnitDB

	row := db.QueryRow("SELECT * FROM unit WHERE label = ?", label)
	if err := row.Scan(&unit.ID, &unit.Label); err != nil {
		if err == sql.ErrNoRows {
			return unit, fmt.Errorf("QueryUnitTableByName %v: no such unit", label)
		}
		return unit, fmt.Errorf("QueryUnitTableByName %v: %v", label, err)
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

func GetAllRecipes(db *sql.DB) ([]Recipe, error) {
	var recs []Recipe

	rows, err := db.Query("SELECT recipe_id FROM recipe")
	if err != nil {
		return nil, fmt.Errorf("GetAllRecipes: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("GetAllRecipes: %v", err)
		}

		rec, err := GetRecipeByID(db, id)
		if err != nil {
			return nil, fmt.Errorf("GetAllRecipes: GetRecipeByID %q: %v", id, err)
		}

		recs = append(recs, rec)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllRecipes: %v", err)
	}

	return recs, nil
}

func SubmitRecipe(db *sql.DB, rec Recipe) (int64, error) {
	recStmt, err := db.Prepare(
		"INSERT OR IGNORE INTO recipe (name, description, instructions) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	ingStmt, err := db.Prepare("INSERT OR IGNORE INTO ingredient (label) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}

	unitStmt, err := db.Prepare("INSERT OR IGNORE INTO unit (label) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}

	riStmt, err := db.Prepare("INSERT OR IGNORE INTO recipe_ingredient (recipe_id, ingredient_id, unit_id, quantity) VALUES (?, ?, ?, ?)")
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
	result, err := recStmt.Exec(rec.Name, rec.Description, strings.Join(rec.Instructions, "|"))
	if err != nil {
		return 0, fmt.Errorf("SubmitRecipe: %v", err)
	}
	recID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("SubmitRecipe: %v", err)
	}

	log.Printf("RECIPE ID: %v\n", recID)

	// ING / UNIT SECTION
	for _, ing := range rec.Ingredients {

		// INGREDIENT; check if exists, if not insert
		var ingID int64
		ingDB, err := QueryIngredientTableByName(db, ing.Label)
		if err != nil {
			result, err = ingStmt.Exec(ing.Label)
			if err != nil {
				return 0, fmt.Errorf("SubmitRecipe: %v", err)
			}
			ingID, err = result.LastInsertId()
			if err != nil {
				return 0, fmt.Errorf("SubmitRecipe: %v", err)
			}
		} else {
			ingID = ingDB.ID
		}

		log.Printf("ING: %v; ID: %v\n", ing.Label, ingID)

		// UNIT; check if exists, if not insert
		var unitID int64
		unitDB, err := QueryUnitTableByName(db, ing.Unit)
		if err != nil {
			result, err = unitStmt.Exec(ing.Unit)
			if err != nil {
				return 0, fmt.Errorf("SubmitRecipe: %v", err)
			}
			unitID, err = result.LastInsertId()
			if err != nil {
				return 0, fmt.Errorf("SubmitRecipe: %v", err)
			}
		} else {
			unitID = unitDB.ID
		}

		log.Printf("UNIT: %v; ID: %v\n", ing.Unit, unitID)

		// RECIPE_INGREDIENT
		result, err = riStmt.Exec(recID, ingID, unitID, ing.Quantity)
		if err != nil {
			return 0, fmt.Errorf("SubmitRecipe: %v", err)
		}
		_, err = result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("SubmitRecipe: %v", err)
		}

		log.Printf("\tri: %v; ing: %v; unit: %v; quantity: %v\n", recID, ingID, unitID, ing.Quantity)
	}

	return recID, nil
}

func GetRecipeByID(db *sql.DB, id int64) (Recipe, error) {
	var rec Recipe

	recDB, err := QueryRecipeTableByID(db, id)
	if err != nil {
		return rec, err
	}

	riDB, err := QueryRecipeIngredientsByID(db, id)
	if err != nil {
		return rec, err
	}

	rec.ID = id
	rec.Name = recDB.Name
	rec.Description = recDB.Description
	rec.Instructions = strings.Split(recDB.Instructions, "|")
	rec.Ingredients = make([]Ingredient, 0)

	for _, ri := range riDB {
		unitDB, err := QueryUnitTableByID(db, ri.UnitID)
		if err != nil {
			return rec, err
		}

		ingDB, err := QueryIngredientTableByID(db, ri.IngredientID)
		if err != nil {
			return rec, err
		}

		rec.Ingredients = append(rec.Ingredients, Ingredient{
			Label:    ingDB.Label,
			Quantity: ri.Quantity,
			Unit:     unitDB.Label,
		})
	}

	return rec, nil
}
