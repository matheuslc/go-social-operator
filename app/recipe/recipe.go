package recipe

// Introduction defines an small text describing the recipe
type Introduction string

// Recipe defines how an recipe is
type Recipe struct {
	Introduction
	Ingredients
	Direction
}
