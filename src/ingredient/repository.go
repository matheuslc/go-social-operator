package ingredient

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	log "github.com/sirupsen/logrus"
)

type Writer interface {
	Create(ingredient Ingredient) error
}

type WriterTransaction interface {
	CreateWithTransaction(transaction neo4j.Transaction, ingredientId string, ingredient Ingredient) error
}

type Repository struct {
	Db neo4j.Driver
	Writer
}

func (repo Repository) CreateWithTransaction(transaction neo4j.Transaction, ingredientId string, ingredientItem Ingredient) error {
	_, err := transaction.Run(
		"CREATE (i:Ingredient { id: $id, food_uuid: $food_uuid, food_type: $food_type, unit_type: $unit_type, amount: $amount })"+
			"RETURN i.id, i.food_uuid, i.food_type, i.unit_type, i.amount",
		map[string]interface{}{
			"id":        ingredientId,
			"food_uuid": ingredientItem.Food().GetID().String(),
			"food_type": ingredientItem.Food().Type(),
			"unit_type": ingredientItem.Unit().Type,
			"amount":    ingredientItem.Unit().Value,
		},
	)

	if err != nil {
		fmt.Println("could not create ingredient", err)
		return err
	}

	fmt.Println("ingredneitn item type", ingredientItem.Food().Type())
	// Ingredient -> Food relantionship
	if ingredientItem.Food().Type() == "animal" {
		log.Info("creating ingredient -> animal food relantionship")
		_, err = transaction.Run(
			"MATCH (i:Ingredient), (f:AnimalFood) WHERE i.id = $ingredient_id AND f.id = $food_id CREATE (i)-[uf:USE_FOOD]->(f)",
			map[string]interface{}{
				"food_id":       ingredientItem.Food().GetID().String(),
				"ingredient_id": ingredientId,
			},
		)

		if err != nil {
			log.Error("could not create ingredient -> animal food relationship", err)
			return err
		}
	} else if ingredientItem.Food().Type() == "recipe" {
		log.Info("creating ingredient -> recipe food relantionship")
		_, err = transaction.Run(
			"MATCH (i:Ingredient), (r:Recipe) WHERE i.id = $ingredient_id AND r.id = $recipe_id CREATE (i)-[uf:USE_FOOD]->(r)",
			map[string]interface{}{
				"ingredient_id": ingredientId,
				"recipe_id":     ingredientItem.Food().GetID().String(),
			},
		)

		if err != nil {
			log.Error("could not create ingredient -> animal relationship", err)
			return err
		}
	} else if ingredientItem.Food().Type() == "plant" {
		log.Info("creating ingredient -> plant food relantionship")
		_, err = transaction.Run(
			"MATCH (i:Ingredient), (f:Food) WHERE i.id = $ingredient_id AND f.id = $food_id CREATE (i)-[uf:USE_FOOD]->(f)",
			map[string]interface{}{
				"food_id":       ingredientItem.Food().GetID().String(),
				"ingredient_id": ingredientId,
			},
		)

		if err != nil {
			log.Error("could not create ingredient -> animal relationship", err)
			return err
		}
	} else if ingredientItem.Food().Type() == "product" {
		log.Info("creating ingredient -> product food relantionship")
		_, err = transaction.Run(
			"MATCH (i:Ingredient), (p:ProductFood) WHERE i.id = $ingredient_id AND p.id = $food_id CREATE (i)-[uf:USE_FOOD]->(p)",
			map[string]interface{}{
				"food_id":       ingredientItem.Food().GetID().String(),
				"ingredient_id": ingredientId,
			},
		)

		if err != nil {
			log.Error("could not create ingredient -> animal relationship", err)
			return err
		}
	} else {
		log.Error("could not create ingredient -> food relationship, unknown food type")
		return err
	}

	return nil
}
