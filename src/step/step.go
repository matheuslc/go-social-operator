package step

import (
	m "github.com/matheuslc/guiomar/src/measurements"
)

// Step defines one step an recipe directions.
type Step interface {
	Description() string
}

type StepPayload struct {
	Description string   `json:"description"`
	Duration    m.Minute `json:"duration"`
	Order       int      `json:"order"`
}

type step struct {
	description string
	duration    m.Minute
	order       int
}

// NewStep creates a new step
func NewStep(description string, duration m.Minute, order int) (Step, error) {
	return &step{
		description,
		duration,
		order,
	}, nil
}

// Description returns the description of the step
func (s step) Description() string {
	return s.description
}
