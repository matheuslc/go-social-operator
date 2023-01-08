package step

import (
	m "github.com/matheuslc/guiomar/src/measurements"
)

// Description is an text used as description of an step
type Description string

// Step defines one step an recipe directions.
type Step interface {
	Description() Description
}

type StepPayload struct {
	Description Description `json:"description"`
	Duration    m.Minute    `json:"duration"`
}

type step struct {
	description Description
	duration    m.Minute
}

// NewStep creates a new step
func NewStep(description Description, duration m.Minute) (Step, error) {
	return &step{
		description,
		duration,
	}, nil
}

// Description returns the description of the step
func (s step) Description() Description {
	return s.description
}
