package food

import "github.com/google/uuid"

// Animal food
type AnimalType string

type Animal struct {
	Id         uuid.UUID `json:"id"`
	Name       `json:"name"`
	AnimalType `json:"type"`
}

func NewAnimal(n Name, t AnimalType) Animal {
	return Animal{
		Name:       n,
		AnimalType: t,
	}
}

func (a Animal) GetID() uuid.UUID {
	return a.Id
}

func (a Animal) Type() string {
	return "animal"
}

func (a Animal) GetName() string {
	return string(a.Name)
}
