package direction

import (
	"github.com/matheuslc/guiomar/src/step"
)

// Direction defines the directions of a recipe
type Direction interface {
	Steps() []step.Step
}

type direction struct {
	steps []step.Step
}

// NewDirection creates a new direction
func NewDirection(steps []step.Step) Direction {
	return &direction{steps}
}

// Steps returns the steps of an recipe directions
func (d direction) Steps() []step.Step {
	return d.steps
}
