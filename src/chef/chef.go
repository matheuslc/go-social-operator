package chef

import (
	"github.com/google/uuid"
)

type Name string
type Role string
type VerifiedEmail string
type UnverifiedEmail error
type Id uuid.UUID

type Email interface {
	Create() UnverifiedEmail
	Verify() VerifiedEmail
}

type Chef struct {
	ID       uuid.UUID `json:"id"`
	Name     `json:"name"`
	Email    VerifiedEmail `json:"email"`
	RoleType Role          `json:"role"`
}

// NewChef its like an helper create a n Chef structure passing its type.
func NewChef(r Role, n Name, email VerifiedEmail) (Chef, error) {
	return Chef{
		Name:     n,
		Email:    VerifiedEmail(email),
		RoleType: r,
	}, nil
}
