package recipe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/direction"
	"github.com/matheuslc/guiomar/src/food"
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
	PreparationTime time.Duration                  `json:"preparation_time"`
	Serving         int64                          `json:"serving"`
	Yield           int64                          `json:"yield"`
}

// NewRecipeHandlerWrapper godoc
// @Summary      Create a new recipe
// @Description  You just need your name and your e-mail
// @Tags         chef
// @Accept       application/json
// @Produce      application/json
// @Param        body body createRecipePayload true "Create a new recipe"
// @Router       /api/recipes [post]
func NewRecipeHandlerWrapper(repo Repository, foodRepository food.Repository, ingredientRepository ingredient.Repository) func(http.ResponseWriter, *http.Request) {
	// there are known race conditions using a closure
	// make sure to test and measure it!
	return func(w http.ResponseWriter, r *http.Request) {
		handler(repo, foodRepository, ingredientRepository, w, r)
	}
}

func handler(repo Repository, foodRepository food.Repository, ingredientRepository ingredient.Repository, w http.ResponseWriter, r *http.Request) {
	payload := createRecipePayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("err", err)
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	ingrs, err := convertIngredients(payload.Ingredients, repo, foodRepository)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not convert ingredients. Params are not the expected")
		return
	}

	stps := make([]step.Step, len(payload.Direction))
	for i, stp := range payload.Direction {
		parsed, err := step.NewStep(stp.Description, stp.Duration)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Could not convert steps. Request params are not the expected")
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
		payload.PreparationTime,
		payload.Serving,
		payload.Yield,
	)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	err = repo.Save(rec, ingredientRepository)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not save recipe")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"ok": "true"})
}

func convertIngredients(payload []ingredient.IngredientPayload, recipeRepository Reader, foodRepository food.Repository) ([]ingredient.Ingredient, error) {
	ingrs := make([]ingredient.Ingredient, len(payload))
	for i, ingr := range payload {
		if ingr.Food.Type == "recipe" {
			f, err := recipeRepository.Find(uuid.MustParse(ingr.Food.ID))
			if err != nil {
				return nil, err
			}

			parsed, err := ingredient.NewIngredient(f, ingr.Amount)
			if err != nil {
				return nil, err
			}

			ingrs[i] = parsed
		} else {
			f, err := foodRepository.Find(uuid.MustParse(ingr.Food.ID), ingr.Food.Type)
			if err != nil {
				return nil, err
			}

			parsed, err := ingredient.NewIngredient(f, ingr.Amount)
			if err != nil {
				return nil, err
			}

			ingrs[i] = parsed
		}
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