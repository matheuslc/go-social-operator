package steps

import (
	"testing"

	food "github.com/matheuslc/guiomar/internal/food"
	ingr "github.com/matheuslc/guiomar/internal/ingredient"
	ingrs "github.com/matheuslc/guiomar/internal/ingredients"
	units "github.com/matheuslc/guiomar/internal/measurements"
	"github.com/matheuslc/guiomar/internal/step"
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
