package recipe

import (
	d "github.com/matheuslc/guiomar/app/direction"
	"github.com/matheuslc/guiomar/app/ingredients"
	ingrs "github.com/matheuslc/guiomar/app/ingredients"
	m "github.com/matheuslc/guiomar/app/measurements"
)

// Introduction defines an small text describing the recipe
type Introduction string

// Recipe defines the structure of a recipe.
type Recipe interface {
	Introduction() Introduction
	Ingredients() ingredients.Ingredients
	Direction() d.Direction
}

// Recipe defines how an recipe is
type recipe struct {
	introduction    Introduction
	ingredients     ingredients.Ingredients
	direction       d.Direction
	cookDuration    m.CookDuration
	preparationTime m.PreparationTime
	serving         m.Serving
	yield           m.Yield
}

// NewRecipe creates a valid recipe
func NewRecipe(
	introduction Introduction,
	ingredients ingrs.Ingredients,
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

func (r recipe) Ingredients() ingredients.Ingredients {
	return r.ingredients
}

func (r recipe) Direction() d.Direction {
	return r.direction
}

func (r recipe) Introduction() Introduction {
	return r.introduction
}
