package recipe

import (
	d "github.com/matheuslc/guiomar/src/direction"
	"github.com/matheuslc/guiomar/src/ingredient"
	m "github.com/matheuslc/guiomar/src/measurements"
)

// Introduction defines an small text describing the recipe
type Introduction string

// Recipe defines the structure of a recipe.
type Recipe interface {
	Introduction() Introduction
	Ingredients() []ingredient.Ingredient
	Direction() d.Direction
}

// Recipe defines how an recipe is
type recipe struct {
	introduction    Introduction
	ingredients     []ingredient.Ingredient
	direction       d.Direction
	cookDuration    m.CookDuration
	preparationTime m.PreparationTime
	serving         m.Serving
	yield           m.Yield
}

// NewRecipe creates a valid recipe
func NewRecipe(
	introduction Introduction,
	ingredients []ingredient.Ingredient,
	direction d.Direction,
	cookDuration m.CookDuration,
	prepartionTime m.PreparationTime,
	serving m.Serving,
	yield m.Yield,
) (Recipe, error) {
	return &recipe{
		introduction:    introduction,
		ingredients:     ingredients,
		direction:       direction,
		cookDuration:    cookDuration,
		preparationTime: prepartionTime,
		serving:         serving,
		yield:           yield,
	}, nil
}

func (r recipe) Ingredients() []ingredient.Ingredient {
	return r.ingredients
}

func (r recipe) Direction() d.Direction {
	return r.direction
}

func (r recipe) Introduction() Introduction {
	return r.introduction
}
