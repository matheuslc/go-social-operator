package food

import (
	"testing"

	"github.com/matheuslc/guiomar/src/measurements"
)

func TestNewFood(t *testing.T) {
	sn := ScientificName("Solanum lycopersicum var. cerasiforme")
	name := Name("Cherry tomato")
	order := Order("Solanales")
	family := Family("Solanaceae")
	genus := Genus("Solanum")
	specie := Specie("S. lycopersicum")
	averageAmount := measurements.UnitType{
		Type:  "gram",
		Value: 1000,
	}

	food := NewVegetalFood(
		sn,
		order,
		family,
		name,
		genus,
		specie,
		averageAmount,
	)

	if food.Name != name {
		t.Errorf("Food name doest has the expected value. Expeced %s, got %s", name, food.Name)
	}

	if food.Genus != genus {
		t.Errorf("Food group doest has the expected value. Expeced %s, got %s", genus, food.Genus)
	}

	if food.Specie != specie {
		t.Errorf("Food subgroup doest has the expected value. Expeced %s, got %s", specie, food.Specie)
	}
}
