package recipe

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/matheuslc/guiomar/src/measurements"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Writer interface {
	Save(r Recipe) (Recipe, error)
}

type Repository struct {
	Db neo4j.Driver
	Writer
}

func (repo Repository) Save(r Recipe) (Recipe, error) {
	session, err := repo.Db.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "neo4j",
	})
	if err != nil {
		fmt.Printf("Cant start a new Neo4j session. Reason: %s", err)
		return Recipe{}, nil
	}

	defer session.Close()

	persistentRecipe, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (r:Recipe {id: $id, summary: $summary, introduction: $introduction, cook_duration: $cook_duration, preparation_time: $preparation_time, serving: $serving, yield: $yield}) "+
				"RETURN r.id, r.summary, r.introduction, r.cook_duration, r.preparation_time, r.serving, r.genus, r.yield",
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
			fmt.Println(result.Record().Values()...)
			return Recipe{
				ID:              uuid.MustParse(result.Record().GetByIndex(0).(string)),
				Introduction:    Introduction(result.Record().GetByIndex(0).(string)),
				Summary:         Summary(result.Record().GetByIndex(1).(string)),
				CookDuration:    measurements.Minute(result.Record().GetByIndex(2).(int64)),
				PreparationTime: measurements.PreparationTime(result.Record().GetByIndex(3).(int64)),
				Serving:         measurements.Serving(result.Record().GetByIndex(4).(int64)),
				Yield:           measurements.Yield(result.Record().GetByIndex(5).(int64)),
			}, nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return Recipe{}, err
	}

	return persistentRecipe.(Recipe), nil
}
