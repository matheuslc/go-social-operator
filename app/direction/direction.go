package direction

import ss "github.com/matheuslc/guiomar/app/steps"

// Direction defines the directions of a recipe
type Direction interface {
	Steps() ss.Steps
}

type direction struct {
	steps ss.Steps
}

// NewDirection creates a new direction
func NewDirection(steps ss.Steps) Direction {
	return &direction{steps}
}

// Steps returns the steps of an recipe directions
func (d direction) Steps() ss.Steps {
	return d.steps
}
