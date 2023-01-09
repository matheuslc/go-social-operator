package ingredient

import (
	"testing"

	f "github.com/matheuslc/guiomar/src/food"
	m "github.com/matheuslc/guiomar/src/measurements"
)

// Scenarios:
// - Convert a unit based on the densitity of the food

func TestNewIngredient(t *testing.T) {
	food := f.NewVegetalFood(
		f.ScientificName("Scientific name"),
		f.Order("order"),
		f.Family("family"),
		f.Name("Cherry Tomato"),
		f.Genus("Vegetables"),
		f.Specie("Fruit vegetables"),
	)

	unit := m.UnitType{
		Type:  "gram",
		Value: 60,
	}

	i, err := NewIngredient(food, unit)

	if err != nil {
		t.Errorf("Cannot create a new ingredient. Error %s", err)
	}

	if i.Food() != food {
		t.Errorf("Ingredient was created with an unexcpeted food.")
	}
}
