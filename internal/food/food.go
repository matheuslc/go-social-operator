package food

import "github.com/google/uuid"

type Uuid uuid.UUID

// Name defines the name of a food. Food names can have variations,
// but they need to have an elected name.
// Cherry Tomato
type Name string

// Specie defines the scientific specie name
type Specie string

// Genus defines the scientific genus name
type Genus string

// Food defines the food struct and its properties
type Food struct {
	Uuid
	Name
	Genus
	Specie
}

// NewFood creates a new food struct with requireds params
func NewFood(n Name, g Genus, sp Specie) Food {
	return Food{
		Name:   n,
		Genus:  g,
		Specie: sp,
	}
}
