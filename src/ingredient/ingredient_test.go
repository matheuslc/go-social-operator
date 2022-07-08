package ingredient

import (
	"testing"

	f "github.com/matheuslc/guiomar/src/food"
	m "github.com/matheuslc/guiomar/src/measurements"
)

func TestNewIngredient(t *testing.T) {
	food := f.NewFood(
		f.ScientificName("solanum lycoperscicum"),
		f.Name("Cherry Tomato"),
		f.Group("Vegetables"),
		f.Subgroup("Fruit vegetables"),
	)

	unit := m.Gram(60)
	i, err := NewIngredient(food, unit)

	if err != nil {
		t.Errorf("Cannot create a new ingredient. Error %s", err)
	}

	if i.Food() != food {
		t.Errorf("Ingredient was created with an unexcpeted food.")
	}

	if i.Unit() != unit {
		t.Errorf("Ingredient was created with an unexcpeted unit. Exptected: %f. Got: %f", unit, i.Unit())
	}
}
