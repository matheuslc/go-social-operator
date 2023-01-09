package food

import "github.com/google/uuid"

// Plant defines the food struct and its properties
type Plant struct {
	Id             uuid.UUID `json:"id"`
	ScientificName `json:"scientific_name"`
	Order          `json:"order"`
	Family         `json:"family"`
	Name           `json:"name"`
	Genus          `json:"genus"`
	Specie         `json:"specie"`
}

// NewFood creates a new food struct with requireds params
func NewVegetalFood(sn ScientificName, o Order, f Family, n Name, g Genus, sp Specie) Plant {
	return Plant{
		ScientificName: sn,
		Order:          o,
		Family:         f,
		Name:           n,
		Genus:          g,
		Specie:         sp,
	}
}

// ID returns the ID of a food
func (f Plant) GetID() uuid.UUID {
	return f.Id
}

func (f Plant) Type() string {
	return "plant"
}

func (f Plant) GetName() string {
	return string(f.Name)
}
