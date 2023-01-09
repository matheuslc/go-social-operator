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
type Amount measurements.Unit

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
