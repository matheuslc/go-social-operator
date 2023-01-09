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
	Id             uuid.UUID
	ScientificName `json:"scientific_name"`
	Order          `json:"order"`
	Family         `json:"family"`
	Name           `json:"name"`
	Genus          `json:"genus"`
	Specie         `json:"specie"`
}

type Animal struct {
	Id         uuid.UUID
	Name       `json:"name"`
	AnimalType `json:"type"`
}

type FoodPublic struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type Fooder interface {
	GetID() uuid.UUID
	Type() string
	GetName() string
}

type FindFoodPayload struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// NewFood creates a new food struct with requireds params
func NewVegetalFood(sn ScientificName, o Order, f Family, n Name, g Genus, sp Specie) Food {
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
func NewAnimal(n Name, t AnimalType) Animal {
	return Animal{
		Name:       n,
		AnimalType: t,
	}
}

// ID returns the ID of a food
func (f Food) GetID() uuid.UUID {
	return f.Id
}

// ID returns the ID of a food
func (a Animal) GetID() uuid.UUID {
	return a.Id
}

func (f Food) Type() string {
	return "vegetal"
}

func (a Animal) Type() string {
	return "animal"
}

func (f Food) GetName() string {
	return string(f.Name)
}

func (a Animal) GetName() string {
	return string(a.Name)
}
