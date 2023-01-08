package food

import (
	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/measurements"
)

// Name defines the name of a food. Food names can have variations,
// but they need to have an elected name.
// Cherry Tomato
type Name string
type ScientificName string
type Specie string
type Genus string
type Order string
type Family string

// Animal food
type AnimalType string

type Amount measurements.Unit

// Food defines the food struct and its properties
type Food struct {
	id             uuid.UUID
	ScientificName `json:"scientific_name"`
	Order          `json:"order"`
	Family         `json:"family"`
	Name           `json:"name"`
	Genus          `json:"genus"`
	Specie         `json:"specie"`
}

type Animal struct {
	id         uuid.UUID
	Name       `json:"name"`
	AnimalType `json:"type"`
}

type Fooder interface {
	ID() uuid.UUID
}

// NewFood creates a new food struct with requireds params
func NewFood(sn ScientificName, o Order, f Family, n Name, g Genus, sp Specie) Food {
	return Food{
		ScientificName: sn,
		Order:          o,
		Family:         f,
		Name:           n,
		Genus:          g,
		Specie:         sp,
	}
}

// NewAnimal creates a new animal struct with requireds params
func NewAnimal(n Name, t AnimalType) (Animal, error) {
	return Animal{
		Name:       n,
		AnimalType: t,
	}, nil
}

// ID returns the ID of a food
func (f Food) ID() uuid.UUID {
	return f.id
}

// ID returns the ID of a food
func (a Animal) ID() uuid.UUID {
	return a.id
}
