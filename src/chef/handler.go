package chef

import (
	"encoding/json"
	"net/http"
)

type createChefPayload struct {
	Role  `json:"role"`
	Name  `json:"name"`
	Email VerifiedEmail `json:"email"`
}

// NewChefHandlerWrapper godoc
// @Summary      Create a new chef
// @Description  You just need your name and your e-mail
// @Tags         chef
// @Accept       application/json
// @Produce      application/json
// @Param        body body createChefPayload true "Create a new chef"
// @Router       /api/chefs [post]
func NewChefHandlerWrapper(repo Repository) func(http.ResponseWriter, *http.Request) {
	// there are known race conditions using a closure
	// make sure to test and measure it!
	return func(w http.ResponseWriter, r *http.Request) {
		handler(repo, w, r)
	}
}

func handler(repo Repository, w http.ResponseWriter, r *http.Request) {
	payload := createChefPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	chef, err := NewChef(
		Role(payload.Role),
		Name(payload.Name),
		VerifiedEmail(payload.Email),
	)

	persistedChef, err := repo.Create(chef)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	respondWithJSON(w, http.StatusOK, persistedChef)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
