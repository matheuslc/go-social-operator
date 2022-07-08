package food

import "testing"

func TestNewFood(t *testing.T) {
	name := Name("Cherry tomato")
	group := Genus("Vegetables")
	subgroup := Specie("Fruit vegetables")

	food := NewFood(
		name,
		group,
		subgroup,
	)
	if food.Name != name {
		t.Errorf("Food name doest has the expected value. Expeced %s, got %s", name, food.Name)
	}

	if food.Genus != group {
		t.Errorf("Food group doest has the expected value. Expeced %s, got %s", group, food.Genus)
	}

	if food.Specie != subgroup {
		t.Errorf("Food subgroup doest has the expected value. Expeced %s, got %s", subgroup, food.Specie)
	}
}
