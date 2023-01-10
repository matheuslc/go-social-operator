package food

import (
	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/measurements"
)

type Product struct {
	Id            uuid.UUID `json:"id"`
	Name          `json:"name"`
	Brand         string `json:"brand"`
	AverageAmount measurements.UnitType
}

func (p Product) GetID() uuid.UUID {
	return p.Id
}

func (p Product) Type() string {
	return "product"
}

func (p Product) GetName() string {
	return string(p.Name)
}

func (p Product) Average() measurements.UnitType {
	return p.AverageAmount
}
