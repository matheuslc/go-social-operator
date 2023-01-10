package ingredient

import (
	"github.com/google/uuid"
	food "github.com/matheuslc/guiomar/src/food"
	units "github.com/matheuslc/guiomar/src/measurements"
)

// Ingredient interface defiens the power of an ingredient
type Ingredient interface {
	GetID() uuid.UUID
	Type() string
	Food() food.Fooder
	Unit() units.UnitType
}

type IngredientPublic struct {
	Food food.FoodPublic `json:"food"`
	Unit units.UnitType  `json:"unit"`
}

type IngredientPayload struct {
	Food   food.FindFoodPayload `json:"food"`
	Amount units.UnitType       `json:"amount"`
}

// ingredient defines an ingredient.
type ingredient struct {
	id   uuid.UUID
	food food.Fooder
	unit units.UnitType
}

// NewIngredient creates a new ingredient
func NewIngredient(f food.Fooder) (Ingredient, error) {
	return ingredient{
		food: f,
		unit: f.Average(),
	}, nil
}

// Food returns the food of an ingredient
func (i ingredient) Food() food.Fooder {
	return i.food
}

// Food returns the food of an ingredient+
func (i ingredient) Unit() units.UnitType {
	return i.unit
}

// Food returns the food of an ingredient+
func (i ingredient) Type() string {
	return i.food.Type()
}

func (i ingredient) GetID() uuid.UUID {
	return i.id
}
