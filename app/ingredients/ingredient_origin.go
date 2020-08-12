package ingredients

// IngredientOrigin describes the origin of the Ingredient
type IngredientOrigin interface {
	New() IngredientOrigin
}

// AnimalOrigin defines the AnimaOrigin option type
type AnimalOrigin string

// VegetableOrigin defines the VegetableOrigin option type
type VegetableOrigin string

// NewAnimalOrigin creates a new type AnimalOrigin
func NewAnimalOrigin() AnimalOrigin {
	return AnimalOrigin("animal")
}

// NewVegetableOrigin a new type VegetableOrigin
func NewVegetableOrigin() VegetableOrigin {
	return VegetableOrigin("vegetable")
}
