package measurements

import "errors"

// Unit represents the base type of all types
type Unit float32

// Minute defines the minute unit type
type Minute Unit

// Serving defines the unit serving
type Serving Unit

// Yield defines the serving people
type Yield Unit

// CookDuration defines the cook duration type
type CookDuration Unit

// PreparationTime defines the time to prepare the recipe
type PreparationTime Unit

type Uniter interface {
	ValueOf() float64
}

var (
	ErrNotConvertible = errors.New("Not convertible")
)

type Conversor interface {
	Convert(to Uniter) (Uniter, error)
}

// UnitToGram creates a new gram with an defined alue
func (unit Unit) UnitToGram() Gram {
	gram := Gram(unit)
	return gram
}

// ValueOf defines
func (unit Unit) ValueOf() float64 {
	return float64(unit)
}
