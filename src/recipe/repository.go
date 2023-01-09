package recipe

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/food"
	"github.com/matheuslc/guiomar/src/ingredient"
	"github.com/matheuslc/guiomar/src/measurements"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	log "github.com/sirupsen/logrus"
)

type Reader interface {
	Find(id uuid.UUID) (Recipe, error)
}

type Writer interface {
	Save(r Recipe) (Recipe, error)
}

type Repository struct {
	Db neo4j.Driver
	Writer
	Reader
}

func (repo Repository) Find(id uuid.UUID) (Recipe, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: "neo4j",
	})

	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return Recipe{}, nil
	}

	defer session.Close()

	recipe, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (r:Recipe)-[i:USE_INGREDIENT]->(ingredients)-[f:USE_FOOD]->(foods) WHERE r.id = $id RETURN r, i, ingredients, foods",
			map[string]interface{}{
				"id": id.String(),
			},
		)

		if err != nil {
			return nil, err
		}

		var rrr Recipe
		ingredientsCollection := []ingredient.Ingredient{}

		for result.Next() {
			r := result.Record().GetByIndex(0).(neo4j.Node)
			ingredientsNode := result.Record().GetByIndex(2).(neo4j.Node)
			foods := result.Record().GetByIndex(3).(neo4j.Node)

			var f food.Fooder
			var u measurements.UnitType

			fmt.Println("type", foods.Props())

			switch foods.Props()["type"] {
			case "animal":
				f = food.Animal{
					Id:   uuid.MustParse(foods.Props()["id"].(string)),
					Name: food.Name(foods.Props()["name"].(string)),
				}
			case "vegetable":
				f = food.Food{
					Id:             uuid.MustParse(foods.Props()["id"].(string)),
					ScientificName: food.ScientificName(foods.Props()["scientific_name"].(string)),
					Order:          food.Order(foods.Props()["order"].(string)),
					Name:           food.Name(foods.Props()["name"].(string)),
					Specie:         food.Specie(foods.Props()["specie"].(string)),
					Family:         food.Family(foods.Props()["family"].(string)),
					Genus:          food.Genus(foods.Props()["genus"].(string)),
				}
			default:
				f = Recipe{
					ID:           uuid.MustParse(foods.Props()["id"].(string)),
					Introduction: Introduction(foods.Props()["introduction"].(string)),
				}
			}

			u = measurements.UnitType{
				Type:  ingredientsNode.Props()["unit_type"].(string),
				Value: ingredientsNode.Props()["amount"].(float64),
			}

			parsedIngredient, err := ingredient.NewIngredient(f, u)
			if err != nil {
				log.Error(err)
				return nil, err
			}

			ingredientsCollection = append(ingredientsCollection, parsedIngredient)

			rrr = Recipe{
				ID:           uuid.MustParse(r.Props()["id"].(string)),
				Summary:      Summary(r.Props()["summary"].(string)),
				Introduction: Introduction(r.Props()["introduction"].(string)),
				CookDuration: measurements.Minute(r.Props()["cook_duration"].(int64)),
				Ingredients:  ingredientsCollection,
			}
		}

		return rrr, nil
	})

	if err != nil {
		return Recipe{}, err
	}

	return recipe.(Recipe), nil
}

func (repo Repository) Save(r Recipe, ingredientRepository ingredient.WriterTransaction) error {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})

	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return err
	}

	defer session.Close()

	_, err = session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (r:Recipe {id: $id, summary: $summary, introduction: $introduction, cook_duration: $cook_duration, preparation_time: $preparation_time, serving: $serving, yield: $yield}) "+
				"RETURN r.id, r.summary, r.introduction, r.cook_duration, r.preparation_time, r.serving, r.yield",
			map[string]interface{}{
				"id":               uuid.New().String(),
				"summary":          r.Summary,
				"introduction":     r.Introduction,
				"cook_duration":    r.CookDuration,
				"preparation_time": r.PreparationTime,
				"serving":          r.Serving,
				"yield":            r.Yield,
			},
		)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			recipeId := uuid.MustParse(result.Record().GetByIndex(0).(string)).String()

			for _, ingredientItem := range r.Ingredients {
				ingredientId := uuid.New().String()

				if err := ingredientRepository.CreateWithTransaction(transaction, ingredientId, ingredientItem); err != nil {
					log.Error("could not create ingredient", err)
					return nil, err
				}

				_, err = transaction.Run(
					"MATCH (r:Recipe), (i:Ingredient) WHERE r.id = $recipe_id AND i.id = $ingredient_id CREATE (r)-[ui:USE_INGREDIENT]->(i)",
					map[string]interface{}{
						"recipe_id":     recipeId,
						"ingredient_id": ingredientId,
					},
				)

				if err != nil {
					log.Error("could not create recipe relantionship with ingredient", err)
					return nil, err
				}
			}

			return nil, nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return err
	}

	return nil
}