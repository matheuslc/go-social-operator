package main

import (
	"fmt"

	units "github.com/matheuslc/guiomar/app/measurements"
	rec "github.com/matheuslc/guiomar/app/recipe"

	food "github.com/matheuslc/guiomar/app/food"
)

func main() {
	food := food.NewFood(
		food.ScientificName("Solanum lycopersicum var. cerasiforme"),
		food.Name("Cherry tomato"),
		food.Group("Vegetables"),
		food.Subgroup("Fruit vegetables"),
	)

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

	fmt.Println(recipe.Introduction)
	fmt.Println(recipe.CookDuration)
	fmt.Println(recipe.Ingredients[0].Food.Name)
}
