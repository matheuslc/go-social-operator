package measurements

import "testing"

func TestGramConversion(t *testing.T) {
	unit := Gram(480)

	converted, err := unit.Convert(Cup(0))
	if err != nil {
		t.Errorf("Cannot convert unit. Error %s", err)
	}

	if converted.ValueOf() != 2 {
		t.Errorf("Expected 2, got %f", converted)
	}
}
