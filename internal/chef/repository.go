package chef

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Writer interface {
	Create(c Chef) (Chef, error)
}

type Repository struct {
	Db neo4j.Driver
}

func (repo Repository) Create(c Chef) (Chef, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return Chef{}, err
	}

	defer session.Close()

	persistedChef, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (chef:Chef {uuid: $uuid, name: $name, email: $email, role: $role}) RETURN chef.uuid, chef.name, chef.email, chef.role",
			map[string]interface{}{
				"uuid":  uuid.New().String(),
				"name":  c.Name,
				"email": c.Email,
				"role":  c.RoleType,
			},
		)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return Chef{
				Uuid:     uuid.MustParse(result.Record().GetByIndex(0).(string)),
				Name:     Name(result.Record().GetByIndex(1).(string)),
				Email:    VerifiedEmail(result.Record().GetByIndex(2).(string)),
				RoleType: Role(result.Record().GetByIndex(3).(string)),
			}, nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return Chef{}, err
	}

	return persistedChef.(Chef), nil
}
