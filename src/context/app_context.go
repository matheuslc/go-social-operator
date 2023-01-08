package context

import (
	"fmt"
	"os"

	"github.com/matheuslc/guiomar/src/chef"
	"github.com/matheuslc/guiomar/src/db"
	"github.com/matheuslc/guiomar/src/food"
	"github.com/matheuslc/guiomar/src/recipe"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type AppContext struct {
	Db               neo4j.Driver
	FoodRepository   food.Repository
	ChefRepository   chef.Repository
	RecipeRepository recipe.Repository
}

// NewAppContext creates a new app context within all dependencies we currently have.
// It's a way to share dependencies instance between different parts of the application.
func NewAppContext() (AppContext, error) {
	db, err := db.New(os.Getenv("NEO4J_HOST"), os.Getenv("NEO4J_USERNAME"), os.Getenv("NEO4J_PASSWORD"))
	if err != nil {
		fmt.Printf("Can't connect to neo4j. Reason: %s", err)
	}

	foodRepository := food.Repository{Db: db}
	chefRepository := chef.Repository{Db: db}
	recipeRepository := recipe.Repository{Db: db}

	return AppContext{
		Db:               db,
		FoodRepository:   foodRepository,
		ChefRepository:   chefRepository,
		RecipeRepository: recipeRepository,
	}, nil
}
