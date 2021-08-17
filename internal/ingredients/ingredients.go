package ingredients

import ing "github.com/matheuslc/guiomar/internal/ingredient"

// ingredients defines an collection of Ingredient
type ingredients []ing.Ingredient

// Ingredients defines what an collection of ingredients can do
type Ingredients interface {
	Add(ing.Ingredient) Ingredients
	First() ing.Ingredient
}

// NewIngredients creates a new ingredient
func NewIngredients() Ingredients {
	is := ingredients{}
	return is
}

// Add adds a new ingredient to the list
func (is ingredients) Add(i ing.Ingredient) Ingredients {
	return append(is, i)
}

// First returns the first element of an slice
func (is ingredients) First() ing.Ingredient {
	return is[0]
}
