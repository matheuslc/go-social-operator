package food

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/measurements"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Reader interface {
	find(id uuid.UUID) (Fooder, error)
}

type Writer interface {
	savePlant(f Plant) (Plant, error)
	saveAnimal(f Animal) (Animal, error)
	saveProduct(p Product) (Product, error)
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
		return Plant{}, nil
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
		case "plant":
			result, err := transaction.Run(
				"MATCH (f:PlantFood {id: $id}) RETURN f.id, f.scientific_name, f.name, f.order, f.family, f.genus, f.specie, f.type",
				map[string]interface{}{
					"id": id.String(),
				},
			)

			if err != nil {
				return nil, err
			}

			if result.Next() {
				return Plant{
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
		case "product":
			result, err := transaction.Run(
				"MATCH (p:ProductFood {id: $id}) RETURN p.id, p.name, p.average_type, p.average_value",
				map[string]interface{}{
					"id": id.String(),
				},
			)

			if err != nil {
				return nil, err
			}

			if result.Next() {
				return Product{
					Id:   uuid.MustParse(result.Record().GetByIndex(0).(string)),
					Name: Name(result.Record().GetByIndex(1).(string)),
					AverageAmount: measurements.UnitType{
						Type:  result.Record().GetByIndex(2).(string),
						Value: result.Record().GetByIndex(3).(float64),
					},
				}, nil
			}

			return nil, result.Err()
		default:
			return nil, errors.New("fatal")
		}
	})

	if err != nil {
		return Plant{}, err
	}

	switch v := (persistedFood).(type) {
	case Plant:
		return v, nil
	case Animal:
		return v, nil
	case Product:
		return v, nil
	default:

		fmt.Println("could not typescast", v)
	}

	return Plant{}, nil
}

func (repo Repository) Save(f Plant) (Plant, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return Plant{}, nil
	}

	defer session.Close()

	persistedFood, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (f:PlantFood {id: $id, scientific_name: $scientific_name, name: $name, order: $order, family: $family, genus: $genus, specie: $specie}) RETURN f.id, f.scientific_name, f.name, f.order, f.family, f.genus, f.specie, f.type",
			map[string]interface{}{
				"id":              uuid.New().String(),
				"scientific_name": f.ScientificName,
				"name":            f.Name,
				"order":           f.Order,
				"family":          f.Family,
				"genus":           f.Genus,
				"specie":          f.Specie,
				"type":            f.Type(),
			},
		)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return Plant{
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
		return Plant{}, err
	}

	return persistedFood.(Plant), nil
}

func (repo Repository) SaveAnimal(f Animal) (Fooder, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return Plant{}, nil
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

func (repo Repository) SaveProduct(p Product) (Fooder, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return Plant{}, nil
	}

	defer session.Close()

	persistedFood, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (f:ProductFood { id: $id, name: $name, type: $type, average_type: $average_type, average_value: $average_value }) RETURN f.id, f.name, f.type, f.average_type, f.average_value",
			map[string]interface{}{
				"id":            uuid.New().String(),
				"name":          p.Name,
				"type":          p.Type(),
				"average_type":  p.AverageAmount.Type,
				"average_value": p.AverageAmount.Value,
			},
		)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return Product{
				Id:   uuid.MustParse(result.Record().GetByIndex(0).(string)),
				Name: Name(result.Record().GetByIndex(1).(string)),
				AverageAmount: measurements.UnitType{
					Type:  result.Record().GetByIndex(3).(string),
					Value: result.Record().GetByIndex(4).(float64),
				},
			}, nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return Product{}, err
	}

	return persistedFood.(Product), nil
}
