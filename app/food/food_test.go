package food

import "testing"

func TestNewFood(t *testing.T) {
	scientificName := ScientificName("Solanum lycopersicum var. cerasiforme")
	name := Name("Cherry tomato")
	group := Group("Vegetables")
	subgroup := Subgroup("Fruit vegetables")

	food := NewFood(
		scientificName,
		name,
		group,
		subgroup,
	)

	if food.ScientificName != scientificName {
		t.Errorf("Food scientificName doest has the expected value. Expeced %s, got %s", scientificName, food.ScientificName)
	}

	if food.Name != name {
		t.Errorf("Food name doest has the expected value. Expeced %s, got %s", name, food.Name)
	}

	if food.Group != group {
		t.Errorf("Food group doest has the expected value. Expeced %s, got %s", group, food.Group)
	}

	if food.Subgroup != subgroup {
		t.Errorf("Food subgroup doest has the expected value. Expeced %s, got %s", subgroup, food.Subgroup)
	}
}
