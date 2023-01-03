package measurements

import "testing"

func TestCupConversion(t *testing.T) {
	unit := Cup(1)
	converted, err := unit.Convert(Gram(0))
	if err != nil {
		t.Errorf("Cannot convert unit. Error %s", err)
	}

	if converted.ValueOf() != 240.0 {
		t.Errorf("Expected 240, got %f", converted)
	}
}
