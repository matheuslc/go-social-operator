package recipe

import (
	"testing"

	"github.com/matheuslc/guiomar/src/direction"

	"github.com/matheuslc/guiomar/src/step"

	"github.com/matheuslc/guiomar/src/food"
	"github.com/matheuslc/guiomar/src/ingredient"
	units "github.com/matheuslc/guiomar/src/measurements"
)

func TestNewRecipe(t *testing.T) {
	foodName := food.Name("Cherry tomato")
	food := food.NewVegetalFood(
		food.ScientificName("Solanum lycopersicum var. cerasiforme"),
		food.Order("Solanales"),
		food.Family("Solanaceae"),
		foodName,
		food.Genus("Fruit vegetables"),
		food.Specie("Cherry tomato"),
	)

	grams := units.UnitType{
		Type:  "gram",
		Value: 60,
	}

	ingr, _ := ingredient.NewIngredient(food, grams)
	collection := []ingredient.Ingredient{}
	ingrs := append(collection, ingr)

	firstStep, err := step.NewStep(
		"Corte em pedaços e adicione ao fogo",
		units.Minute(2),
		0,
	)
	if err != nil {
		t.Errorf("Error when creating a new step. Error: %s", err)
	}

	lastStep, err := step.NewStep(
		"Adicione o molho na geladeira",
		units.Minute(2),
		1,
	)

	if err != nil {
		t.Errorf("Error when creating a new step. Error: %s", err)
	}

	stepsCollection := []step.Step{}
	stepsCollection = append(stepsCollection, firstStep)
	stepsCollection = append(stepsCollection, lastStep)

	d, err := direction.NewDirection(stepsCollection)
	if err != nil {
		t.Errorf("Error when creating a new direction. Error: %s", err)
	}

	intro := Introduction("Receita de teste")
	summary := Summary("Resumo da receita de teste")

	rec, err := NewRecipe(
		summary,
		intro,
		ingrs,
		d,
		50,
		20,
		6,
		3,
	)

	if err != nil {
		t.Errorf("Error when creating a new recipe. Error: %s", err)
	}

	if rec.Introduction != intro {
		t.Errorf("Introduction was not the expected. Expected: %s. Got: %s", intro, rec.Introduction)
	}
}
