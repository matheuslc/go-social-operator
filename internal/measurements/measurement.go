package measurements

// Unit represents the base type of all types
type Unit float32

// Mensurable defines the type tha can be used as a mensurable
// e.g 1.0
type Mensurable interface {
	ValueOf() float64
}

// Microgram defines the gram unit type
// e.g 1.0
type Microgram Unit

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

// UnitToGram creates a new gram with an defined alue
func (unit Unit) UnitToGram() Gram {
	gram := Gram(unit)
	return gram
}

// ValueOf defines
func (unit Unit) ValueOf() float64 {
	return float64(unit)
}

// ValueOf definmes
func (gram Gram) ValueOf() float64 {
	return float64(gram)
}
