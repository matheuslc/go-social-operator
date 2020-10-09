package recipe

// ingredients defines an collection of Ingredient
type ingredients []Ingredient

// Ingredients defines what an collection of ingredients can do
type Ingredients interface {
	Add(Ingredient) ingredients
	First() Ingredient
}

// NewIngredients creates a new ingredient
func NewIngredients() Ingredients {
	is := ingredients{}
	return is
}

// Add adds a new ingredient to the list
func (is ingredients) Add(i Ingredient) ingredients {
	return append(is, i)
}

// First returns the first element of an slice
func (is ingredients) First() Ingredient {
	return is[0]
}
