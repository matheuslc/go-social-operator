package food

import (
	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/measurements"
)

// Animal food
type AnimalType string

type Animal struct {
	Id            uuid.UUID `json:"id"`
	Name          `json:"name"`
	AnimalType    `json:"type"`
	AverageAmount measurements.UnitType `json:"average_amount"`
}

func NewAnimal(n Name, t AnimalType, av measurements.UnitType) Animal {
	return Animal{
		Name:          n,
		AnimalType:    t,
		AverageAmount: av,
	}
}

func (a Animal) GetID() uuid.UUID {
	return a.Id
}

func (a Animal) Type() string {
	return string(FoodTypeAnimal)
}

func (a Animal) GetName() string {
	return string(a.Name)
}

func (a Animal) Average() measurements.UnitType {
	return a.AverageAmount
}