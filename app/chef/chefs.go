package chef

import (
	"errors"

	"github.com/google/uuid"
)

type name string
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
	PublicEmail() verifiedEmail
	NewChef(name, verifiedEmail) (Chef, error)
}

type amauter struct {
	ID    id
	Name  name
	Email verifiedEmail
	Type  string
}

type professional struct {
	ID    id
	Name  name
	Email verifiedEmail
	Type  string
}

// NewChef its like an helper create an Chef structure passing its type.
func NewChef(t string, name name, email string) (Chef, error) {
	if t == "professional" {
		email := verifiedEmail(email)
		chef := professional{}
		return chef.NewChef(name, email)
	}

	return nil, errors.New("Error when creating an chef")
}

// NewChef create a new chef. It validates it's fields too
func (p professional) NewChef(name name, email verifiedEmail) (Chef, error) {
	return professional{
		Name:  name,
		Email: email,
		Type:  "professional",
	}, nil
}

// PublicEmail return the amauter email
func (p professional) PublicEmail() verifiedEmail {
	return p.Email
}

// NewChef creates a new amauter chef
func (a amauter) NewChef(name name, email verifiedEmail) (Chef, error) {
	return amauter{
		Name:  name,
		Email: email,
		Type:  "amauter",
	}, nil
}

// PublicEmail return the amauter email
func (a amauter) PublicEmail() verifiedEmail {
	return a.Email
}
