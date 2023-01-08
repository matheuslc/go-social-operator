package recipe

import (
	"encoding/json"
	"net/http"

	"github.com/matheuslc/guiomar/src/direction"
	"github.com/matheuslc/guiomar/src/ingredient"
	m "github.com/matheuslc/guiomar/src/measurements"
)

type createRecipePayload struct {
	Summary         Summary                 `json:"summary"`
	Introduction    Introduction            `json:"introduction"`
	CookDuration    m.Minute                `json:"cook_duration"`
	Ingredients     []ingredient.Ingredient `json:"ingredients"`
	Directions      []direction.Direction   `json:"directions"`
	PreparationTime m.PreparationTime       `json:"preparation_time"`
	Serving         m.Serving               `json:"serving"`
	Yield           m.Yield                 `json:"yield"`
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

	rec, err := NewRecipe(
		Summary(payload.Summary),
		Introduction(payload.Introduction),
		payload.Ingredients,
		payload.Directions,
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
