package food

import (
	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/measurements"
)

// Plant defines the food struct and its properties
type Plant struct {
	Id             uuid.UUID `json:"id"`
	ScientificName `json:"scientific_name"`
	Order          `json:"order"`
	Family         `json:"family"`
	Name           `json:"name"`
	Genus          `json:"genus"`
	Specie         `json:"specie"`
	AverageAmount  measurements.UnitType `json:"average_amount"`
}

// NewFood creates a new food struct with requireds params
func NewVegetalFood(sn ScientificName, o Order, f Family, n Name, g Genus, sp Specie, av measurements.UnitType) Plant {
	return Plant{
		ScientificName: sn,
		Order:          o,
		Family:         f,
		Name:           n,
		Genus:          g,
		Specie:         sp,
		AverageAmount:  av,
	}
}

// ID returns the ID of a food
func (f Plant) GetID() uuid.UUID {
	return f.Id
}

func (f Plant) Type() string {
	return string(FoodTypePlant)
}

func (f Plant) GetName() string {
	return string(f.Name)
}

func (f Plant) Average() measurements.UnitType {
	return f.AverageAmount
}
