package recipe

import (
	"testing"

	"github.com/matheuslc/guiomar/app/direction"

	"github.com/matheuslc/guiomar/app/step"
	"github.com/matheuslc/guiomar/app/steps"

	"github.com/matheuslc/guiomar/app/food"
	"github.com/matheuslc/guiomar/app/ingredient"
	"github.com/matheuslc/guiomar/app/ingredients"
	units "github.com/matheuslc/guiomar/app/measurements"
)

func TestNewRecipe(t *testing.T) {
	foodName := food.Name("Cherry tomato")
	food := food.NewFood(
		food.ScientificName("Solanum lycopersicum var. cerasiforme"),
		foodName,
		food.Group("Vegetables"),
		food.Subgroup("Fruit vegetables"),
	)
	grams := units.Gram(10)
	ingr, _ := ingredient.NewIngredient(food, grams)
	collection := ingredients.NewIngredients()
	ingrs := collection.Add(ingr)

	firstStep := step.NewStep(
		step.Description("Corte em pedaços e adicione ao fogo"),
		units.Minute(2),
		ingrs,
	)

	lastStep := step.NewStep(
		step.Description("Adicione o molho na geladeira"),
		units.Minute(2),
		ingrs,
	)

	stepsCollection := steps.NewSteps()
	stepsCollection.Add(firstStep)
	stepsCollection.Add(lastStep)
	directions := direction.NewDirection(stepsCollection)

	intro := Introduction("Receita de teste")

	rec, err := NewRecipe(
		intro,
		ingrs,
		directions,
		units.CookDuration(50),
		units.PreparationTime(20),
		units.Serving(6),
		units.Yield(3),
	)

	if err != nil {
		t.Errorf("Error when creating a new recipe. Error: %s", err)
	}

	if rec.Introduction() != intro {
		t.Errorf("Introduction was not the expected. Expected: %s. Got: %s", intro, rec.Introduction())
	}
}
