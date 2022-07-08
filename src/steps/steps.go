package steps

import (
	"github.com/matheuslc/guiomar/src/step"
)

// Steps defines a collection of steps.
type Steps interface {
	Add(step.Step) Steps
	First() step.Step
}

type steps []step.Step

// NewSteps creates a new steps collection
func NewSteps() Steps {
	st := &steps{}
	return st
}

// Add adds a new ingredient to the list
func (s steps) Add(st step.Step) Steps {
	return append(s, st)
}

// First returns the first step of an steps collection
func (s steps) First() step.Step {
	return s[0]
}
