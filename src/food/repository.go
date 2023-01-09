package food

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Reader interface {
	find(id uuid.UUID) (Food, error)
}

type Writer interface {
	save(f Food) (Food, error)
	saveAnimal(f Animal) (Fooder, error)
}

type Repository struct {
	Db neo4j.Driver
	Writer
	Reader
}

func (repo Repository) Find(id uuid.UUID, foodType string) (Fooder, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return Food{}, nil
	}

	defer session.Close()

	persistedFood, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		switch foodType {
		case "animal":
			result, err := transaction.Run(
				"MATCH (f:AnimalFood {id: $id}) RETURN f.id, f.name, f.type",
				map[string]interface{}{
					"id": id.String(),
				},
			)

			if err != nil {
				return nil, err
			}

			if result.Next() {
				return Animal{
					Id:         uuid.MustParse(result.Record().GetByIndex(0).(string)),
					Name:       Name(result.Record().GetByIndex(1).(string)),
					AnimalType: AnimalType(result.Record().GetByIndex(2).(string)),
				}, nil
			}

			return nil, result.Err()
		default:
			result, err := transaction.Run(
				"MATCH (f:Food {id: $id}) RETURN f.id, f.scientific_name, f.name, f.order, f.family, f.genus, f.specie, f.type",
				map[string]interface{}{
					"id": id.String(),
				},
			)

			if err != nil {
				return nil, err
			}

			if result.Next() {
				return Food{
					Id:             uuid.MustParse(result.Record().GetByIndex(0).(string)),
					ScientificName: ScientificName(result.Record().GetByIndex(1).(string)),
					Name:           Name(result.Record().GetByIndex(2).(string)),
					Order:          Order(result.Record().GetByIndex(3).(string)),
					Family:         Family(result.Record().GetByIndex(4).(string)),
					Genus:          Genus(result.Record().GetByIndex(5).(string)),
					Specie:         Specie(result.Record().GetByIndex(6).(string)),
				}, nil
			}

			return nil, result.Err()
		}
	})

	if err != nil {
		return Food{}, err
	}

	switch v := (persistedFood).(type) {
	case Food:
		return v, nil
	case Animal:
		return v, nil
	}

	return Food{}, nil
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
			return Food{
				Id:             uuid.MustParse(result.Record().GetByIndex(0).(string)),
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

func (repo Repository) SaveAnimal(f Animal) (Fooder, error) {
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
			"CREATE (f:AnimalFood { id: $id, name: $name, type: $type }) RETURN f.id, f.name, f.type",
			map[string]interface{}{
				"id":   uuid.New().String(),
				"name": f.Name,
				"type": f.Type(),
			},
		)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return Animal{
				Id:         uuid.MustParse(result.Record().GetByIndex(0).(string)),
				Name:       Name(result.Record().GetByIndex(1).(string)),
				AnimalType: AnimalType(result.Record().GetByIndex(2).(string)),
			}, nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return Animal{}, err
	}

	return persistedFood.(Animal), nil
}
