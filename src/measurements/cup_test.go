package measurements

import "testing"

func TestCupConversion(t *testing.T) {
	toGram := UnitType{
		Type:  "gram",
		Value: 240,
	}

	cup := Cup(1)

	converted, err := cup.Convert(toGram)
	if err != nil {
		t.Errorf("Cannot convert unit. Error %s", err)
	}

	if converted.ValueOf() != 240 {
		t.Errorf("Expected 240, got %f", converted.ValueOf())
	}
}
