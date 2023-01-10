package category

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Writer interface {
	Create(c Category) (Category, error)
}

type Reader interface {
	Find(id uuid.UUID) (Category, error)
}

type Repository struct {
	Db neo4j.Driver
}

func (repo Repository) Create(c Category) (Category, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return Category{}, err
	}

	defer session.Close()

	persistedCategory, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (c:Category {id: $id, name: $name}) RETURN c.id, c.name",
			map[string]interface{}{
				"id":   uuid.New().String(),
				"name": c.Name,
			},
		)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return Category{
				ID:   uuid.MustParse(result.Record().GetByIndex(0).(string)),
				Name: result.Record().GetByIndex(1).(string),
			}, nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return Category{}, err
	}

	return persistedCategory.(Category), nil
}

func (repo Repository) Find(id uuid.UUID) (Category, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return Category{}, err
	}

	defer session.Close()

	persistedCategory, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"MATCH (c:Category {id: $id) RETURN c.id, c.name",
			map[string]interface{}{
				"id": id.String(),
			},
		)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return Category{
				ID:   uuid.MustParse(result.Record().GetByIndex(0).(string)),
				Name: result.Record().GetByIndex(1).(string),
			}, nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return Category{}, err
	}

	return persistedCategory.(Category), nil
}
