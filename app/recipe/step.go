package recipe

import m "github.com/matheuslc/guiomar/app/measurements"

// Description is an text used as description of an step
type Description string

// Step defines one step an recipe directions.
type Step struct {
	Description
	Duration m.Minute
	ingredients
}

// Steps defines a collection of steps.
type Steps []Step
