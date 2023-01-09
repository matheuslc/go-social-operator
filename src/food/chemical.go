package food

import "github.com/google/uuid"

type Chemical struct {
	Id                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	ExternalReference string    `json:"external_reference"`
}
