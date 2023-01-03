package ingredient

import (
	food "github.com/matheuslc/guiomar/src/food"
	units "github.com/matheuslc/guiomar/src/measurements"
)

// Ingredient interface defiens the power of an ingredient
type Ingredient interface {
	Food() food.Food
	Unit() units.Conversor
}

// ingredient defines an ingredient.
type ingredient struct {
	food food.Food
	unit units.Conversor
}

// NewIngredient creates a new ingredient
func NewIngredient(f food.Food, u units.Conversor) (Ingredient, error) {
	return ingredient{
		food: f,
		unit: u,
	}, nil
}

// Food returns the food of an ingredient
func (i ingredient) Food() food.Food {
	return i.food
}

// Food returns the food of an ingredient+
func (i ingredient) Unit() units.Conversor {
	return i.unit
}
