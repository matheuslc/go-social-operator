package steps

import (
	"testing"

	food "github.com/matheuslc/guiomar/src/food"
	ingr "github.com/matheuslc/guiomar/src/ingredient"
	ingrs "github.com/matheuslc/guiomar/src/ingredients"
	units "github.com/matheuslc/guiomar/src/measurements"
	"github.com/matheuslc/guiomar/src/step"
)

func TestNewSteps(t *testing.T) {
	food := food.NewFood(
		food.ScientificName("Solanum lycopersicum var. cerasiforme"),
		food.Name("Cherry tomato"),
		food.Group("Vegetables"),
		food.Subgroup("Fruit vegetables"),
	)

	grams := units.Gram(10)
	i, _ := ingr.NewIngredient(food, grams)
	ingredients := ingrs.NewIngredients()
	ingredients.Add(i)

	stepDescription := step.Description("Cozinhe esse tomate")
	firstStep := step.NewStep(
		stepDescription,
		units.Minute(10),
		ingredients,
	)

	stepsCollection := NewSteps()
	added := stepsCollection.Add(firstStep)

	if added.First().Description() != stepDescription {
		t.Errorf("Step description was not the expected. The step was not added correctly")
	}
}
