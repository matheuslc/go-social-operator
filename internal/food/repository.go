package food

import (
	"fmt"

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

func (repo Repository) Save(f Food) (string, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return "", nil
	}

	defer session.Close()

	persistedFood, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (f:Food) SET f.name = $name, f.genus = $genus, f.specie RETURN f",
			map[string]interface{}{
				"name":   f.Name,
				"genus":  f.Genus,
				"specie": f.Specie,
			},
		)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return "", err
	}

	return persistedFood.(string), nil
}
