package chef

import (
	"errors"

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

// Chef interface returns an chef which can be an amauter or a professional
type Chef interface {
	PublicEmail() VerifiedEmail
}

// Amauter defines and specialiazed type of chef
type Amauter struct {
	Name
	Email    VerifiedEmail
	RoleType Role
}

// Professional defines an specialized type of chef
type Professional struct {
	Name
	Email    VerifiedEmail
	RoleType Role
}

// NewChef its like an helper create an Chef structure passing its type.
func NewChef(r Role, n Name, email string) (Chef, error) {
	if r == "professional" {
		email := VerifiedEmail(email)
		chef := Professional{
			Name:     n,
			Email:    email,
			RoleType: r,
		}

		return chef, nil
	}

	if r == "amauter" {
		email := VerifiedEmail(email)
		chef := Amauter{
			Name:     n,
			Email:    email,
			RoleType: r,
		}

		return chef, nil
	}

	return nil, errors.New("error when creating a chef")
}

// PublicEmail return the amauter email
func (p Professional) PublicEmail() VerifiedEmail {
	return p.Email
}

// PublicEmail return the amauter email
func (a Amauter) PublicEmail() VerifiedEmail {
	return a.Email
}
