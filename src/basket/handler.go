package basket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/food"
	"github.com/matheuslc/guiomar/src/ingredient"
	"github.com/matheuslc/guiomar/src/measurements"
	"github.com/matheuslc/guiomar/src/recipe"

	log "github.com/sirupsen/logrus"
)

type basketPayload struct {
	Recipes []string `json:"recipes"`
}

// NewBasketHandlerWerapper godoc
// @Summary      Create a new basket based on many recipes
// @Description  List the recipes you want
// @Tags         basket
// @Accept       application/json
// @Produce      application/json
// @Param        body body basketPayload true "create a new basket based on some payload"
// @Router       /api/basket [post]
func NewBasketHandlerWerapper(recipeRepository recipe.Reader) func(http.ResponseWriter, *http.Request) {
	// there are known race conditions using a closure
	// make sure to test and measure it!
	return func(w http.ResponseWriter, r *http.Request) {
		handler(recipeRepository, w, r)
	}
}

func handler(recipeRepository recipe.Reader, w http.ResponseWriter, r *http.Request) {
	payload := basketPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		fmt.Println("err", err)
		respondWithError(w, http.StatusBadRequest, "Request params are not the expected")
		return
	}

	recipeCollection := make([]recipe.Recipe, len(payload.Recipes))

	for index, recipeID := range payload.Recipes {
		persitedRecipe, err := recipeRepository.Find(uuid.MustParse(recipeID))
		if err != nil {
			log.Errorf("could find out the recipe. Err: %v", err)
			respondWithError(w, http.StatusBadRequest, "Could not find recipe")
			return
		}

		recipeCollection[index] = persitedRecipe
	}

	ingredientCollection := []ingredient.IngredientPublic{}
	for _, recipe := range recipeCollection {
		for _, ing := range recipe.Ingredients {
			ingredientPublic := ingredient.IngredientPublic{
				Food: food.FoodPublic{
					ID:            ing.Food().GetID().String(),
					Name:          ing.Food().GetName(),
					Type:          ing.Food().Type(),
					AverageAmount: ing.Food().Average(),
				},
				Unit: measurements.UnitType{
					Type:  ing.Unit().Type,
					Value: ing.Unit().Value,
				},
			}

			ingredientCollection = append(ingredientCollection, ingredientPublic)
		}
	}

	// Data map to public structures for rest response
	recipesPublicCollection := []recipe.RecipePublic{}
	for _, r := range recipeCollection {
		innerIngredientCollection := []ingredient.IngredientPublic{}
		for _, ing := range r.Ingredients {
			ingredientPublic := ingredient.IngredientPublic{
				Food: food.FoodPublic{
					ID:            ing.Food().GetID().String(),
					Name:          ing.Food().GetName(),
					Type:          ing.Food().Type(),
					AverageAmount: ing.Food().Average(),
				},
				Unit: measurements.UnitType{
					Type:  ing.Unit().Type,
					Value: ing.Unit().Value,
				},
			}

			innerIngredientCollection = append(innerIngredientCollection, ingredientPublic)
		}

		recipesPublicCollection = append(recipesPublicCollection, recipe.RecipePublic{
			ID:            r.ID,
			Introduction:  string(r.Introduction),
			Ingredients:   innerIngredientCollection,
			Summary:       string(r.Summary),
			Serving:       r.Serving,
			Yield:         r.Yield,
			Category:      r.Category,
			Direction:     r.Direction,
			CookDuration:  time.Duration(r.CookDuration),
			AverageAmount: r.AverageAmount,
		})
	}

	unifyIngredients := make(map[string]ingredient.IngredientPublic)
	unifiedCollection := []ingredient.IngredientPublic{}
	for _, ing := range ingredientCollection {
		key := ing.Food.ID
		if _, ok := unifyIngredients[key]; ok {
			result := unifyIngredients[key].Unit.Value + ing.Unit.Value
			ing.Unit.Value = result

			unifyIngredients[key] = ing
		} else {
			unifyIngredients[key] = ing
		}
	}

	for _, ing := range unifyIngredients {
		unifiedCollection = append(unifiedCollection, ing)
	}

	bsk := BasketPublic{Recipes: recipesPublicCollection, Ingredients: unifiedCollection}

	respondWithJSON(w, http.StatusOK, bsk)
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
