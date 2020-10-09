package recipe

import m "github.com/matheuslc/guiomar/app/measurements"

// Introduction defines an small text describing the recipe
type Introduction string

// Recipe defines how an recipe is
type Recipe struct {
	Introduction
	ingredients
	Direction
	m.CookDuration
	m.PreparationTime
	m.Serving
	m.Yield
}
