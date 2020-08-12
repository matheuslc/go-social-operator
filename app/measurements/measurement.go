package measurements

// Unit represents the base type of all types
type Unit float32

// Gram defines the gram unit type
// e.g 1.0
type Gram Unit

// Minute defines the minute unit type
type Minute Unit

// Centimeter defines the centimer unit type
// e.g 2.0
type Centimeter Unit

// Serving defines the unit serving
type Serving Unit

// Yield defines the serving people
type Yield Unit

// CookDuration defines the cook duration type
type CookDuration Unit

// PreparationTime defines the time to prepare the recipe
type PreparationTime Unit

// NewGram creates a new gram with an defined alue
func NewGram(amount float64) Gram {
	unit := Gram(amount)
	return unit
}
