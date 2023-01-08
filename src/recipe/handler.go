package recipe

import (
	"encoding/json"
	"net/http"

	"github.com/matheuslc/guiomar/src/direction"
	"github.com/matheuslc/guiomar/src/ingredient"
	m "github.com/matheuslc/guiomar/src/measurements"
	"github.com/matheuslc/guiomar/src/step"
)

type createRecipePayload struct {
	Summary         Summary                        `json:"summary"`
	Introduction    Introduction                   `json:"introduction"`
	CookDuration    m.Minute                       `json:"cook_duration"`
	Ingredients     []ingredient.IngredientPayload `json:"ingredients"`
	Direction       []step.StepPayload             `json:"directions"`
	PreparationTime m.PreparationTime              `json:"preparation_time"`
	Serving         m.Serving                      `json:"serving"`
	Yield           m.Yield                        `json:"yield"`
}

// NewRecipeHandlerWrapper godoc
// @Summary      Create a new recipe
// @Description  You just need your name and your e-mail
// @Tags         chef
// @Accept       application/json
// @Produce      application/json
// @Param        body body createRecipePayload true "Create a new recipe"
// @Router       /api/recipes [post]
func NewRecipeHandlerWrapper(repo Repository) func(http.ResponseWriter, *http.Request) {
	// there are known race conditions using a closure
	// make sure to test and measure it!
	return func(w http.ResponseWriter, r *http.Request) {
		handler(repo, w, r)
	}
}

func handler(repo Repository, w http.ResponseWriter, r *http.Request) {
	payload := createRecipePayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	ingrs, err := convertIngredients(payload.Ingredients)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	stps := make([]step.Step, len(payload.Direction))
	for i, stp := range payload.Direction {
		parsed, err := step.NewStep(stp.Description, stp.Duration)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
			return
		}

		stps[i] = parsed
	}

	di, err := direction.NewDirection(stps)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	rec, err := NewRecipe(
		Summary(payload.Summary),
		Introduction(payload.Introduction),
		ingrs,
		di,
		m.Minute((payload.CookDuration)),
		m.PreparationTime(payload.PreparationTime),
		m.Serving(payload.Serving),
		m.Yield(payload.Yield),
	)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	persistedRecipe, err := repo.Save(rec)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	respondWithJSON(w, http.StatusOK, persistedRecipe)
}

func convertIngredients(payload []ingredient.IngredientPayload) ([]ingredient.Ingredient, error) {
	ingrs := make([]ingredient.Ingredient, len(payload))
	for i, ingr := range payload {
		parsed, err := ingredient.NewIngredient(ingr.Food, ingr.Amount)
		if err != nil {
			return nil, err
		}

		ingrs[i] = parsed
	}

	return ingrs, nil
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err := w.Write(response)
	if err != nil {
		panic(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
