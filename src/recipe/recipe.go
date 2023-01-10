package recipe

import (
	"time"

	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/category"
	"github.com/matheuslc/guiomar/src/direction"
	"github.com/matheuslc/guiomar/src/ingredient"
	"github.com/matheuslc/guiomar/src/measurements"
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
	Direction       direction.Direction     `json:"direction"`
	Category        category.Category       `json:"category"`
	CookDuration    m.Minute                `json:"cook_duration"`
	PreparationTime time.Duration           `json:"preparation_time"`
	Serving         int64                   `json:"serving"`
	Yield           int64                   `json:"yield"`
	AverageAmount   measurements.UnitType   `json:"average_amount"`
}

type RecipePublic struct {
	ID            uuid.UUID                     `json:"id"`
	Introduction  string                        `json:"introduction"`
	Summary       string                        `json:"summary"`
	Ingredients   []ingredient.IngredientPublic `json:"ingredients"`
	Direction     direction.Direction           `json:"direction"`
	Category      category.Category             `json:"category"`
	CookDuration  time.Duration                 `json:"cook_duration"`
	Serving       int64                         `json:"serving"`
	Yield         int64                         `json:"yield"`
	AverageAmount measurements.UnitType         `json:"average_amount"`
}

// NewRecipe creates a valid recipe
func NewRecipe(
	summary Summary,
	introduction Introduction,
	ingredients []ingredient.Ingredient,
	cat category.Category,
	direction direction.Direction,
	cookDuration m.Minute,
	prepartionTime time.Duration,
	serving int64,
	yield int64,
	averageAmount measurements.UnitType,
) (Recipe, error) {
	return Recipe{
		Summary:         summary,
		Introduction:    introduction,
		Ingredients:     ingredients,
		Direction:       direction,
		Category:        cat,
		CookDuration:    cookDuration,
		PreparationTime: prepartionTime,
		Serving:         serving,
		Yield:           yield,
		AverageAmount:   averageAmount,
	}, nil
}

func (r Recipe) SummaryHTML() string {
	return string(r.Summary)
}

func (r Recipe) GetID() uuid.UUID {
	return r.ID
}

func (r Recipe) Type() string {
	return "recipe"
}

func (r Recipe) GetName() string {
	return string(r.Introduction)
}

func (r Recipe) Average() measurements.UnitType {
	return r.AverageAmount
}
