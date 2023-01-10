package category

import (
	"encoding/json"
	"net/http"
)

// NewChefHandlerWrapper godoc
// @Summary      Create a new category
// @Description  You just need your name and your e-mail
// @Tags         chef
// @Accept       application/json
// @Produce      application/json
// @Param        body body CreateCategoryPayload true "Create a new category"
// @Router       /api/category [post]
func NewCategorHandlerWrapper(repo Writer) func(http.ResponseWriter, *http.Request) {
	// there are known race conditions using a closure
	// make sure to test and measure it!
	return func(w http.ResponseWriter, r *http.Request) {
		handler(repo, w, r)
	}
}

func handler(repo Writer, w http.ResponseWriter, r *http.Request) {
	payload := CreateCategoryPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	categoryToPersist := Category{Name: payload.Name}
	cat, err := repo.Create(categoryToPersist)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create the category")
		return
	}

	respondWithJSON(w, http.StatusOK, cat)
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
