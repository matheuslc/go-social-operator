package ingredient

import (
	food "github.com/matheuslc/guiomar/src/food"
	units "github.com/matheuslc/guiomar/src/measurements"
)

// Ingredient interface defiens the power of an ingredient
type Ingredient interface {
	Food() food.Fooder
	Unit() units.Conversor
}

type IngredientPayload struct {
	Food   food.Fooder     `json:"food"`
	Amount units.Conversor `json:"amount"`
}

// ingredient defines an ingredient.
type ingredient struct {
	food food.Fooder
	unit units.Conversor
}

// NewIngredient creates a new ingredient
func NewIngredient(f food.Fooder, u units.Conversor) (Ingredient, error) {
	return ingredient{
		food: f,
		unit: u,
	}, nil
}

// Food returns the food of an ingredient
func (i ingredient) Food() food.Fooder {
	return i.food
}

// Food returns the food of an ingredient+
func (i ingredient) Unit() units.Conversor {
	return i.unit
}
