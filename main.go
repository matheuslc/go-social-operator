package main

import (
	"fmt"
	"log"
	"net/http"

	food "github.com/matheuslc/guiomar/app/food"
	units "github.com/matheuslc/guiomar/app/measurements"
	rec "github.com/matheuslc/guiomar/app/recipe"
)

type CustomType float64

type product struct {
	Name  string
	Price float64
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	food := food.NewFood(
		food.ScientificName("Solanum lycopersicum var. cerasiforme"),
		food.Name("Cherry tomato"),
		food.Group("Vegetables"),
		food.Subgroup("Fruit vegetables"),
	)

	a := CustomType(4)
	b := CustomType(4)

	if a == b {
		fmt.Println("true")
	}

	grams := units.Gram(10)

	ingredient := rec.Ingredient{
		Food: food,
		Unit: grams,
	}

	ingredients := rec.Ingredients{ingredient}

	firstStep := rec.Step{
		Description: rec.Description("Cozinhe essa tomate"),
		Duration:    units.Minute(10),
		Ingredients: ingredients,
	}

	steps := rec.Steps{firstStep}
	direction := rec.Direction{Steps: steps}

	recipe := rec.Recipe{
		Introduction:    rec.Introduction("Essa receita é show"),
		Ingredients:     ingredients,
		Direction:       direction,
		CookDuration:    units.CookDuration(10.0),
		PreparationTime: units.PreparationTime(20.0),
		Serving:         units.Serving(6.0),
		Yield:           units.Yield(10.0),
	}

	fmt.Println(recipe.Ingredients[0].Food.Name)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
