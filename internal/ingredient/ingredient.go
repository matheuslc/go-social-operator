package ingredient

import (
	food "github.com/matheuslc/guiomar/internal/food"
	units "github.com/matheuslc/guiomar/internal/measurements"
)

// Ingredient interface defiens the power of an ingredient
type Ingredient interface {
	Food() food.Food
	Unit() units.Mensurable
}

// ingredient defines an ingredient.
type ingredient struct {
	food food.Food
	unit units.Mensurable
}

// NewIngredient creates a new ingredient
func NewIngredient(f food.Food, u units.Mensurable) (Ingredient, error) {
	return ingredient{
		food: f,
		unit: u,
	}, nil
}

// Food returns the food of an ingredient
func (i ingredient) Food() food.Food {
	return i.food
}

// Food returns the food of an ingredient
func (i ingredient) Unit() units.Mensurable {
	return i.unit
}
