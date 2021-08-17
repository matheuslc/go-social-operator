package tool

import (
	"github.com/google/uuid"
)

type Name string
type UUID uuid.UUID

type Appliance struct {
	ID   UUID
	Name Name
}

type Utensil struct {
	ID   UUID
	Name Name
}
