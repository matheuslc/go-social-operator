package ingredients

import (
	"testing"

	food "github.com/matheuslc/guiomar/app/food"
	ing "github.com/matheuslc/guiomar/app/ingredient"
	units "github.com/matheuslc/guiomar/app/measurements"
)

func TestNewIngredients(t *testing.T) {
	foodName := food.Name("Cherry tomato")
	food := food.NewFood(
		food.ScientificName("Solanum lycopersicum var. cerasiforme"),
		foodName,
		food.Group("Vegetables"),
		food.Subgroup("Fruit vegetables"),
	)

	grams := units.Gram(10)
	ingr, _ := ing.NewIngredient(food, grams)

	collection := NewIngredients()
	ingrs := collection.Add(ingr)

	if ingrs.First().Food().Name != foodName {
		t.Errorf("Food name was not the expected")
	}
}
