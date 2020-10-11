package recipe

import (
	"github.com/matheuslc/guiomar/app/ingredients"
	m "github.com/matheuslc/guiomar/app/measurements"
)

// Introduction defines an small text describing the recipe
type Introduction string

// Recipe defines how an recipe is
type Recipe struct {
	Introduction
	ingredients.Ingredients
	Direction
	m.CookDuration
	m.PreparationTime
	m.Serving
	m.Yield
}
