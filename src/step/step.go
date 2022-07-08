package step

import (
	ingrs "github.com/matheuslc/guiomar/src/ingredients"
	m "github.com/matheuslc/guiomar/src/measurements"
)

// Description is an text used as description of an step
type Description string

// Step defines one step an recipe directions.
type Step interface {
	Description() Description
}

type step struct {
	description Description
	duration    m.Minute
	ingredients ingrs.Ingredients
}

// NewStep creates a new step
func NewStep(description Description, duration m.Minute, ingredients ingrs.Ingredients) Step {
	return &step{
		description,
		duration,
		ingredients,
	}
}

// Description returns the description of the step
func (s step) Description() Description {
	return s.description
}
