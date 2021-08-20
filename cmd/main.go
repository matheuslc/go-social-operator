package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/matheuslc/guiomar/internal/db"
	"github.com/matheuslc/guiomar/internal/direction"
	food "github.com/matheuslc/guiomar/internal/food"
	ingr "github.com/matheuslc/guiomar/internal/ingredient"
	ingrs "github.com/matheuslc/guiomar/internal/ingredients"
	units "github.com/matheuslc/guiomar/internal/measurements"
	rec "github.com/matheuslc/guiomar/internal/recipe"
	"github.com/matheuslc/guiomar/internal/step"
	"github.com/matheuslc/guiomar/internal/steps"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type AppContext struct {
	Db             neo4j.Driver
	FoodRepository food.Repository
}

func NewAppContext() (AppContext, error) {
	db, err := db.New(os.Getenv("NEO4J_HOST"), os.Getenv("NEO4J_USERNAME"), os.Getenv("NEO4J_PASSWORD"))
	if err != nil {
		fmt.Printf("Can't connect to neo4j. Reason: %s", err)
	}

	foodRepository := food.Repository{
		Db: db,
	}

	return AppContext{
		Db:             db,
		FoodRepository: foodRepository,
	}, nil
}

func (app AppContext) start() error {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	app, _ := NewAppContext()

	food := food.NewFood(
		food.ScientificName("Solanum lycopersicum var. cerasiforme"),
		food.Name("Cherry tomato"),
		food.Group("Vegetables"),
		food.Subgroup("Fruit vegetables"),
	)

	persistedFood, err := app.FoodRepository.Save(food)

	if err != nil {
		fmt.Println("Error while persisting. Reason: %s", err)
	} else {
		fmt.Println("Food F F F F %s", persistedFood)
	}

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

	app.start()
}
