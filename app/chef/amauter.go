package chef

import "github.com/google/uuid"

type name string
type verifieldEmail string
type unverifiedEmail string
type id uuid.UUID

type Email interface {
	Create() unverifiedEmail
	Verify() verifieldEmail
}

type amauter struct {
	ID    id
	Name  name
	Email VerifiedEmail
}

type professional struct {
	ID    id
	Name  name
	Email verifieldEmail
}
