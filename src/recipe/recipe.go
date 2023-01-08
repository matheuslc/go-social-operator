package recipe

import (
	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/direction"
	"github.com/matheuslc/guiomar/src/ingredient"
	m "github.com/matheuslc/guiomar/src/measurements"
)

// Recipe defines the structure of a recipe.
type Reciper interface {
	SummaryHTML() string
}

type Introduction string
type Summary string

// Recipe defines how an recipe is
type Recipe struct {
	ID              uuid.UUID
	Introduction    Introduction            `json:"introduction"`
	Summary         Summary                 `json:"summary"`
	Ingredients     []ingredient.Ingredient `json:"ingredients"`
	Directions      []direction.Direction   `json:"direction"`
	CookDuration    m.Minute                `json:"cook_duration"`
	PreparationTime m.PreparationTime       `json:"preparation_time"`
	Serving         m.Serving               `json:"serving"`
	Yield           m.Yield                 `json:"yield"`
}

// NewRecipe creates a valid recipe
func NewRecipe(
	summary Summary,
	introduction Introduction,
	ingredients []ingredient.Ingredient,
	directions []direction.Direction,
	cookDuration m.Minute,
	prepartionTime m.PreparationTime,
	serving m.Serving,
	yield m.Yield,
) (Recipe, error) {
	return Recipe{
		Summary:         summary,
		Introduction:    introduction,
		Ingredients:     ingredients,
		Directions:      directions,
		CookDuration:    cookDuration,
		PreparationTime: prepartionTime,
		Serving:         serving,
		Yield:           yield,
	}, nil
}

func (r Recipe) SummaryHTML() string {
	return string(r.Summary)
}
