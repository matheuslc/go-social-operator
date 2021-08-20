package chef

import (
	"github.com/google/uuid"
)

type Name string
type Role string
type VerifiedEmail string
type UnverifiedEmail error
type Id uuid.UUID

// Email interface defines what an Email can do.
type Email interface {
	Create() UnverifiedEmail
	Verify() VerifiedEmail
}

// Amauter defines and specialiazed type of chef
type Chef struct {
	Uuid uuid.UUID
	Name
	Email    VerifiedEmail
	RoleType Role
}

// NewChef its like an helper create an Chef structure passing its type.
func NewChef(r Role, n Name, email string) (Chef, error) {
	return Chef{
		Name:     n,
		Email:    VerifiedEmail(email),
		RoleType: r,
	}, nil
}
