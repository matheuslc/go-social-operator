package recipe

import (
	food "github.com/matheuslc/guiomar/app/food"
	units "github.com/matheuslc/guiomar/app/measurements"
)

// Ingredient defines an ingredient.
type Ingredient struct {
	food.Food
	units.Gram
}

// Ingredients defines an collection of Ingredient
type Ingredients []Ingredient
