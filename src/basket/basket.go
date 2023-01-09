package basket

import (
	"github.com/matheuslc/guiomar/src/ingredient"
	"github.com/matheuslc/guiomar/src/recipe"
)

type BasketPublic struct {
	Recipes     []recipe.RecipePublic         `json:"recipes"`
	Ingredients []ingredient.IngredientPublic `json:"ingredients"`
}

type Basket struct {
	Recipes     []recipe.Recipe         `json:"recipes"`
	Ingredients []ingredient.Ingredient `json:"ingredients"`
}
