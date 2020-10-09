package chef

import (
	"errors"

	"github.com/google/uuid"
)

type name string
type role string
type verifiedEmail string
type unverifiedEmail error
type id uuid.UUID

// Email interface defines what an Email can do.
type Email interface {
	Create() unverifiedEmail
	Verify() verifiedEmail
}

// Chef interface returns an chef which can be an amauter or a professional
type Chef interface {
	Name() name
	NewChef(name, verifiedEmail) (Chef, error)
	PublicEmail() verifiedEmail
	Role() role
}

type amauter struct {
	id
	name
	email    verifiedEmail
	roleType role
}

type professional struct {
	id
	name
	email    verifiedEmail
	roleType role
}

// NewChef its like an helper create an Chef structure passing its type.
func NewChef(r role, n name, email string) (Chef, error) {
	if r == "professional" {
		email := verifiedEmail(email)
		chef := professional{}
		return chef.NewChef(n, email)
	}

	return nil, errors.New("Error when creating a chef")
}

// NewChef create a new chef. It validates it's fields too
func (p professional) NewChef(name name, email verifiedEmail) (Chef, error) {
	return professional{
		name:     name,
		email:    email,
		roleType: role("professional"),
	}, nil
}

// Name returns the professional chef name
func (p professional) Name() name {
	return p.name
}

// PublicEmail return the amauter email
func (p professional) PublicEmail() verifiedEmail {
	return p.email
}

// Role returns the professional chef name
func (p professional) Role() role {
	return p.roleType
}

// NewChef creates a new amauter chef
func (a amauter) NewChef(name name, email verifiedEmail) (Chef, error) {
	return amauter{
		name:     name,
		email:    email,
		roleType: role("amauter"),
	}, nil
}

// Name returns the amauter chef name
func (a amauter) Name() name {
	return a.name
}

// PublicEmail return the amauter email
func (a amauter) PublicEmail() verifiedEmail {
	return a.email
}

// Role returns the professional chef name
func (a amauter) Role() role {
	return a.roleType
}
