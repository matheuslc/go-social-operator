package food

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Reader interface {
}

type Writer interface {
	save(f Food) (Food, error)
}

type Repository struct {
	Db neo4j.Driver
	Writer
	Reader
}

func (repo Repository) Save(f Food) (Food, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return Food{}, nil
	}

	defer session.Close()

	persistedFood, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (f:Food {id: $id, scientific_name: $scientific_name, name: $name, order: $order, family: $family, genus: $genus, specie: $specie}) RETURN f.id, f.scientific_name, f.name, f.order, f.family, f.genus, f.specie",
			map[string]interface{}{
				"id":              uuid.New().String(),
				"scientific_name": f.ScientificName,
				"name":            f.Name,
				"order":           f.Order,
				"family":          f.Family,
				"genus":           f.Genus,
				"specie":          f.Specie,
			},
		)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			fmt.Println(result.Record().Values()...)
			return Food{
				ID:             uuid.MustParse(result.Record().GetByIndex(0).(string)),
				ScientificName: ScientificName(result.Record().GetByIndex(1).(string)),
				Name:           Name(result.Record().GetByIndex(2).(string)),
				Order:          Order(result.Record().GetByIndex(3).(string)),
				Family:         Family(result.Record().GetByIndex(4).(string)),
				Genus:          Genus(result.Record().GetByIndex(5).(string)),
				Specie:         Specie(result.Record().GetByIndex(6).(string)),
			}, nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return Food{}, err
	}

	return persistedFood.(Food), nil
}
