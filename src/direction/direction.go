package direction

import (
	"github.com/matheuslc/guiomar/src/step"
)

// Direction defines the directions of a recipe
type Direction interface {
	Steps() []step.Step
}

type DirectionPayload struct {
	Steps []step.StepPayload `json:"steps"`
}

type direction struct {
	steps []step.Step
}

// NewDirection creates a new direction
func NewDirection(steps []step.Step) (Direction, error) {
	return &direction{steps}, nil
}

// Steps returns the steps of an recipe directions
func (d direction) Steps() []step.Step {
	return d.steps
}
