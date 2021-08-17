package food

// ScientificName defines the scientific name of a food
// e.g. Solanum lycopersicum var. cerasiforme (cherry tomato)
type ScientificName string

// Name defines the name of a food. Food names can have variations,
// but they need to have an elected name.
// Cherry Tomato
type Name string

// Group defines the group of the food.
// e.g Vegetables
type Group string

// Subgroup defines the sub-group of the food.
// e.g Fruit vegetables
type Subgroup string

// Food defines the food struct and its properties
type Food struct {
	ScientificName
	Name
	Group
	Subgroup
}

// NewFood creates a new food struct with requireds params
func NewFood(s ScientificName, n Name, g Group, sb Subgroup) Food {
	return Food{
		s,
		n,
		g,
		sb,
	}
}
