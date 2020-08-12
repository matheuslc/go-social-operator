package ingredients

import (
	"testing"
)

func TestNewAnimalOrigin(t *testing.T) {
	origin := NewAnimalOrigin()

	if origin != "animal" {
		t.Errorf("Expect an animal type, got %s", origin)
	}
}

func TestNewVegetableOrigin(t *testing.T) {
	origin := NewVegetableOrigin()

	if origin != "vegetable" {
		t.Errorf("Expect an vegetable type, got %s", origin)
	}
}
