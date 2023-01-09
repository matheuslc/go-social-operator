package food

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matheuslc/guiomar/src/measurements"
)

type createFoodPayload struct {
	ScientificName ScientificName        `json:"scientific_name"`
	Order          Order                 `json:"order"`
	Family         Family                `json:"family"`
	Name           Name                  `json:"name"`
	Genus          Genus                 `json:"genus"`
	Specie         Specie                `json:"specie"`
	Type           string                `json:"type"`
	AnimalType     string                `json:"animal_type"`
	AverageAmount  measurements.UnitType `json:"average_amount"`
}

// NewFoodHandlerWrapper godoc
// @Summary      Create a new food
// @Description  Creates a new food which can be used within recipes
// @Tags         food
// @Accept       application/json
// @Produce      application/json
// @Param        body body createFoodPayload true "Create a new food"
// @Router       /api/foods [post]
func NewFoodHandlerWrapper(repo Repository) func(http.ResponseWriter, *http.Request) {
	// there are known race conditions using a closure
	// make sure to test and measure it!
	return func(w http.ResponseWriter, r *http.Request) {
		handler(repo, w, r)
	}
}

func handler(repo Repository, w http.ResponseWriter, r *http.Request) {
	var err error
	payload := createFoodPayload{}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	var food Fooder
	var persistedFood Fooder

	switch payload.Type {
	case "animal":
		food = NewAnimal(payload.Name, AnimalType(payload.AnimalType))
		persistedFood, err = repo.SaveAnimal(food.(Animal))
	case "plant":
		food = NewVegetalFood(
			ScientificName(payload.ScientificName),
			Order(payload.Order),
			Family(payload.Family),
			Name(payload.Name),
			Genus(payload.Genus),
			Specie(payload.Specie),
		)
		persistedFood, err = repo.Save(food.(Food))
	case "product":
		food = Product{
			Name:          payload.Name,
			AverageAmount: payload.AverageAmount,
		}

		persistedFood, err = repo.SaveProduct(food.(Product))
	default:
		fmt.Println("Invalid food type")
	}

	if err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	respondWithJSON(w, http.StatusOK, persistedFood)
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
