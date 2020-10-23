package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matheuslc/guiomar/app/direction"
	food "github.com/matheuslc/guiomar/app/food"
	ingr "github.com/matheuslc/guiomar/app/ingredient"
	ingrs "github.com/matheuslc/guiomar/app/ingredients"
	units "github.com/matheuslc/guiomar/app/measurements"
	rec "github.com/matheuslc/guiomar/app/recipe"
	"github.com/matheuslc/guiomar/app/step"
	"github.com/matheuslc/guiomar/app/steps"
)

type product struct {
	Name  string
	Price float64
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
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

	firstStep := step.NewStep(
		step.Description("Cozinhe esse tomate"),
		units.Minute(10),
		ingredients,
	)

	stepsCollection := steps.NewSteps()
	added := stepsCollection.Add(firstStep)
	direction := direction.NewDirection(added)

	recipe, _ := rec.NewRecipe(
		rec.Introduction("Essa receita é show"),
		ingredients,
		direction,
		units.CookDuration(10.0),
		units.PreparationTime(20.0),
		units.Serving(6.0),
		units.Yield(10.0),
	)

	fmt.Println(recipe.Direction().Steps().First().Description())

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
