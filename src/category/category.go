package category

import "github.com/google/uuid"

type SetCategoryPayload struct {
	ID string `json:"id"`
}

type CreateCategoryPayload struct {
	Name string `json:"name"`
}

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
