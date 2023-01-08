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

// Specie defines the scientific specie name
type Specie string
type Genus string

type Order string

type Family string

type Amount measurements.Unit

// Food defines the food struct and its properties
type Food struct {
	ID             uuid.UUID `json:"id"`
	ScientificName `json:"scientific_name"`
	Order          `json:"order"`
	Family         `json:"family"`
	Name           `json:"name"`
	Genus          `json:"genus"`
	Specie         `json:"specie"`
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
