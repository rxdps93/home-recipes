package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: MAJOR - use mutex for all db access operations

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

type RecipeDB struct {
	ID           int64  `json:"recipe_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
	Source       string `json:"source"`
}

type IngredientDB struct {
	ID    int64  `json:"ingredient_id"`
	Label string `json:"label"`
}

type UnitDB struct {
	ID    int64  `json:"unit_id"`
	Label string `json:"label"`
}

type TagDB struct {
	ID    int64  `json:"tag_id"`
	Label string `json:"label"`
}

type RecipeIngredientDB struct {
	ID           int64   `json:"rec_ing_id"`
	RecipeID     int64   `json:"recipe_id"`
	IngredientID int64   `json:"ingredient_id"`
	UnitID       int64   `json:"unit_id"`
	Quantity     float64 `json:"quantity"`
}

type RecipeTagDB struct {
	ID       int64 `json:"rec_tag_id"`
	RecipeID int64 `json:"recipe_id"`
	TagID    int64 `json:"tag_id"`
}

var db *sql.DB = nil

func Connect() {
	con, err := sql.Open("sqlite3", "recipes.db")
	if err != nil {
		log.Fatal(err)
	}
	db = con
}

func Disconnect() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
	db = nil
}

func RemoveRecipeByID(id int64) {
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
func WipeDatabase() {
	db.Exec("DELETE FROM recipe")
	db.Exec("DELETE FROM ingredient")
	db.Exec("DELETE FROM unit")
	db.Exec("DELETE FROM recipe_ingredient")
	db.Exec("DELETE FROM tag")
	db.Exec("DELETE FROM recipe_tag")
}

func AddTestChocolateMilkshakeRecipe() (int64, error) {
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
		Tags: []string{
			"Dessert",
			"Drink",
			"Comfort Food",
			"Family Recipe",
		},
		Source: "My Mom",
	}

	id, err := SubmitRecipe(rec)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func AddTestFreshGuacamoleRecipe() (int64, error) {
	rec := Recipe{
		Name:        "Fresh Guacamole",
		Description: "Delicious guacamole - great on chips!",
		Instructions: []string{
			"Cut the avocados in half, open them, and place the green innards in a bowl.",
			"Mash the avocados.",
			"Add lime juice; continue to mix.",
			"Chop the cilantro, tomato, and onion; Add them to the avocado.",
			"Add salt to taste and mix thoroughly.",
		},
		Ingredients: []Ingredient{
			{
				Label:    "Avocado",
				Quantity: 4,
				Unit:     "",
			},
			{
				Label:    "Red Onion",
				Quantity: 0.5,
				Unit:     "",
			},
			{
				Label:    "Tomato",
				Quantity: 2,
				Unit:     "",
			},
			{
				Label:    "Cilantro",
				Quantity: 2,
				Unit:     "Tablespoons",
			},
			{
				Label:    "Lime",
				Quantity: 2,
				Unit:     "",
			},
		},
		Tags: []string{
			"Dip",
			"Mexican",
			"From Website",
		},
		Source: "https://based.cooking/guacamole/",
	}

	id, err := SubmitRecipe(rec)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func AddTestGarlicLimeChicken() (int64, error) {
	rec := Recipe{
		Name:        "Garlic Lime Chicken",
		Description: "An easy freezer meal recipe that can come together quickly in a slow cooker or in the oven.",
		Instructions: []string{
			"Add all of the ingredients into a large gallon sized ziplock freezer bag.",
			"Freeze until needed or allow to marinate in the refrigerator overnight.",
			"If preparing in the oven: Pour contents of bag into an oven safe dish and bake at 400° F for 25-35 minutes, or until an internal temperature of 165° F.",
			"If preparing in the slow cooker: Pour contents of bag into slow clooker and let cook for 4-5 hours on high.",
		},
		Ingredients: []Ingredient{
			{
				Label:    "Chicken Breast",
				Quantity: 2,
				Unit:     "Lbs",
			},
			{
				Label:    "Onion",
				Quantity: 1,
				Unit:     "",
			},
			{
				Label:    "Garlic",
				Quantity: 4,
				Unit:     "Cloves",
			},
			{
				Label:    "Lime",
				Quantity: 1,
				Unit:     "",
			},
			{
				Label:    "Cilantro",
				Quantity: 0.25,
				Unit:     "Cup",
			},
			{
				Label:    "Chili Powder",
				Quantity: 1,
				Unit:     "Teaspoon",
			},
			{
				Label:    "Chicken Stock",
				Quantity: 2,
				Unit:     "Cups",
			},
			{
				Label:    "Red Pepper Flakes",
				Quantity: 0.5,
				Unit:     "Teaspoons",
			},
		},
		Tags: []string{
			"Freezer Meals",
			"Dinner",
			"From Website",
		},
		Source: "https://www.primallyinspired.com/garlic-lime-chicken-recipe-paleo/",
	}

	id, err := SubmitRecipe(rec)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func AddTestGrilledCheeseRecipe() (int64, error) {
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
		Tags: []string{
			"Sandwich",
			"Lunch",
			"Comfort Food",
			"Family Recipe",
		},
		Source: "Me",
	}

	id, err := SubmitRecipe(rec)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func queryRecipeTableByID(id int64) (RecipeDB, error) {
	var rec RecipeDB

	row := db.QueryRow("SELECT * FROM recipe WHERE recipe_id = ?", id)
	if err := row.Scan(&rec.ID, &rec.Name, &rec.Description, &rec.Instructions, &rec.Source); err != nil {
		if err == sql.ErrNoRows {
			return rec, fmt.Errorf("queryRecipeTableByID %d: no such recipe", id)
		}
		return rec, fmt.Errorf("queryRecipeTableByID %d: %v", id, err)
	}
	return rec, nil
}

func queryIngredientTableByID(id int64) (IngredientDB, error) {
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

func queryIngredientTableByName(label string) (IngredientDB, error) {
	var ing IngredientDB

	row := db.QueryRow("SELECT * FROM ingredient WHERE label = ?", label)
	if err := row.Scan(&ing.ID, &ing.Label); err != nil {
		if err == sql.ErrNoRows {
			return ing, fmt.Errorf("queryIngredientTableByName %v: no such ingredient", label)
		}
		return ing, fmt.Errorf("queryIngredientTableByName %v: %v", label, err)
	}
	return ing, nil
}

func queryUnitTableByID(id int64) (UnitDB, error) {
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

func queryUnitTableByName(label string) (UnitDB, error) {
	var unit UnitDB

	row := db.QueryRow("SELECT * FROM unit WHERE label = ?", label)
	if err := row.Scan(&unit.ID, &unit.Label); err != nil {
		if err == sql.ErrNoRows {
			return unit, fmt.Errorf("queryUnitTableByName %v: no such unit", label)
		}
		return unit, fmt.Errorf("queryUnitTableByName %v: %v", label, err)
	}
	return unit, nil
}

func queryTagTableByID(id int64) (TagDB, error) {
	var tag TagDB
	row := db.QueryRow("SELECT * FROM tag WHERE tag_id = ?", id)
	if err := row.Scan(&tag.ID, &tag.Label); err != nil {
		if err == sql.ErrNoRows {
			return tag, fmt.Errorf("queryTagTableByID %d: no such tag", id)
		}
		return tag, fmt.Errorf("queryTagTableByID: %d: %v", id, err)
	}
	return tag, nil
}

func queryTagTableByName(label string) (TagDB, error) {
	var tag TagDB
	row := db.QueryRow("SELECT * FROM tag WHERE label = ?", label)
	if err := row.Scan(&tag.ID, &tag.Label); err != nil {
		if err == sql.ErrNoRows {
			return tag, fmt.Errorf("queryTagTableByName: %v: no such tag", label)
		}
		return tag, fmt.Errorf("queryTagTableByName %v: %v", label, err)
	}
	return tag, nil
}

func queryRecipeIngredientsByRecipeID(recID int64) ([]RecipeIngredientDB, error) {
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

func queryRecipeTagByRecipeID(recID int64) ([]RecipeTagDB, error) {
	var tags []RecipeTagDB

	rows, err := db.Query("SELECT * FROM recipe_tag WHERE recipe_id = ?", recID)
	if err != nil {
		return nil, fmt.Errorf("queryRecipeTagByRecipeID %q: %v", recID, err)
	}
	defer rows.Close()
	for rows.Next() {
		var rt RecipeTagDB
		if err := rows.Scan(&rt.ID, &rt.RecipeID, &rt.TagID); err != nil {
			return nil, fmt.Errorf("queryRecipeTagByRecipeID %q: %v", recID, err)
		}
		tags = append(tags, rt)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryRecipeTagByRecipeID %q: %v", recID, err)
	}

	return tags, nil
}

func queryRecipeTagByTagID(tagID int64) ([]RecipeTagDB, error) {
	var tags []RecipeTagDB

	rows, err := db.Query("SELECT * FROM recipe_tag WHERE tag_id = ?", tagID)
	if err != nil {
		return nil, fmt.Errorf("queryRecipeTagByTagID %q: %v", tagID, err)
	}
	defer rows.Close()
	for rows.Next() {
		var rt RecipeTagDB
		if err := rows.Scan(&rt.ID, &rt.RecipeID, &rt.TagID); err != nil {
			return nil, fmt.Errorf("queryRecipeTagByTagID %q: %v", tagID, err)
		}
		tags = append(tags, rt)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queryRecipeTagByTagID %q: %v", tagID, err)
	}

	return tags, nil
}

func GetAllTags() ([]string, error) {
	var tags []string

	rows, err := db.Query("SELECT label FROM tag")
	if err != nil {
		return nil, fmt.Errorf("GetAllTags: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, fmt.Errorf("GetAllTags: %v", err)
		}

		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllTags: %v", err)
	}

	return tags, nil
}

func GetRecipesFiltered(nameQuery string, tagQuery []string) ([]Recipe, error) {
	var recs []Recipe

	baseQuery := `SELECT recipe.recipe_id from recipe
    JOIN recipe_tag on recipe_tag.recipe_id = recipe.recipe_id
    JOIN tag on recipe_tag.tag_id = tag.tag_id
    WHERE name like '%' || ? || '%'`

	group := "\ngroup by recipe.recipe_id\n"
	args := []any{nameQuery}
	if len(tagQuery) == 0 {
		baseQuery += group
	} else {
		tagPart := " and\ntag.label in ("
		for _, tag := range tagQuery {
			tagPart += "?"
			if tag != tagQuery[len(tagQuery)-1] {
				tagPart += ", "
			}

			args = append(args, tag)
		}
		tagPart += ")"

		baseQuery += tagPart + group
		baseQuery += "having count(distinct tag.label) = ?"

		args = append(args, len(tagQuery))
	}

	log.Printf("baseQuery: %v", baseQuery)
	log.Printf("arguments: %v", args)
	log.Printf("nameQuery: %v", nameQuery)

	stmt, err := db.Prepare(baseQuery)
	if err != nil {
		return nil, fmt.Errorf("GetRecipesFiltered prepare: %v", err)
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("GetRecipesFiltered query: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("GetRecipesFiltered scan: %v", err)
		}

		rec, err := GetRecipeByID(id)
		if err != nil {
			return nil, fmt.Errorf("GetRecipesFiltered: GetRecipeByID %q: %v", id, err)
		}

		recs = append(recs, rec)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetRecipesFiltered err: %v", err)
	}

	return recs, nil
}

func GetAllRecipes() ([]Recipe, error) {
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

		rec, err := GetRecipeByID(id)
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

func SubmitRecipe(rec Recipe) (int64, error) {
	recStmt, err := db.Prepare(
		"INSERT OR IGNORE INTO recipe (name, description, instructions, source) VALUES (?, ?, ?, ?)")
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

	tagStmt, err := db.Prepare("INSERT OR IGNORE INTO tag (label) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}

	rtStmt, err := db.Prepare("INSERT OR IGNORE INTO recipe_tag (recipe_id, tag_id) VALUES (?, ?)")
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
	 * Repeat similar process for tags
	 */
	// RECIPE SECTION
	result, err := recStmt.Exec(rec.Name, rec.Description, strings.Join(rec.Instructions, "|"), rec.Source)
	if err != nil {
		return 0, fmt.Errorf("SubmitRecipe: %v", err)
	}
	recID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("SubmitRecipe: %v", err)
	}

	// TAG SECTION
	for _, tag := range rec.Tags {
		// TAG; check if exists, if not insert
		var tagID int64
		tagDB, err := queryTagTableByName(tag)
		if err != nil {
			result, err = tagStmt.Exec(tag)
			if err != nil {
				return 0, fmt.Errorf("SubmitRecipe: %v", err)
			}
			tagID, err = result.LastInsertId()
			if err != nil {
				return 0, fmt.Errorf("SubmitRecipe: %v", err)
			}
		} else {
			tagID = tagDB.ID
		}

		result, err = rtStmt.Exec(recID, tagID)
		if err != nil {
			return 0, fmt.Errorf("SubmitRecipe: %v", err)
		}
		_, err = result.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("SubmitRecipe: %v", err)
		}
	}

	// ING / UNIT SECTION
	for _, ing := range rec.Ingredients {

		// INGREDIENT; check if exists, if not insert
		var ingID int64
		ingDB, err := queryIngredientTableByName(ing.Label)
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

		// UNIT; check if exists, if not insert
		var unitID int64
		unitDB, err := queryUnitTableByName(ing.Unit)
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

		// RECIPE_INGREDIENT
		result, err = riStmt.Exec(recID, ingID, unitID, ing.Quantity)
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

func GetRecipeByID(id int64) (Recipe, error) {
	var rec Recipe

	recDB, err := queryRecipeTableByID(id)
	if err != nil {
		return rec, err
	}

	riDB, err := queryRecipeIngredientsByRecipeID(id)
	if err != nil {
		return rec, err
	}

	rtDB, err := queryRecipeTagByRecipeID(id)
	if err != nil {
		return rec, err
	}

	rec.ID = id
	rec.Name = recDB.Name
	rec.Description = recDB.Description
	rec.Instructions = strings.Split(recDB.Instructions, "|")
	rec.Ingredients = make([]Ingredient, 0)
	rec.Tags = make([]string, 0)
	rec.Source = recDB.Source

	for _, ri := range riDB {
		unitDB, err := queryUnitTableByID(ri.UnitID)
		if err != nil {
			return rec, err
		}

		ingDB, err := queryIngredientTableByID(ri.IngredientID)
		if err != nil {
			return rec, err
		}

		rec.Ingredients = append(rec.Ingredients, Ingredient{
			Label:    ingDB.Label,
			Quantity: ri.Quantity,
			Unit:     unitDB.Label,
		})
	}

	for _, rt := range rtDB {
		tagDB, err := queryTagTableByID(rt.TagID)
		if err != nil {
			return rec, err
		}

		rec.Tags = append(rec.Tags, tagDB.Label)
	}

	return rec, nil
}

func GetAllRecipesForTagName(tag string) ([]Recipe, error) {
	var recs []Recipe

	tagDB, err := queryTagTableByName(tag)
	if err != nil {
		return nil, err
	}

	rts, err := queryRecipeTagByTagID(tagDB.ID)
	if err != nil {
		return nil, err
	}

	for _, rt := range rts {
		rec, err := GetRecipeByID(rt.RecipeID)
		if err != nil {
			return nil, err
		}

		recs = append(recs, rec)
	}

	return recs, nil
}
