package measurements

import "testing"

func TestGramConversion(t *testing.T) {
	gram := Gram(480)

	toCup := UnitType{
		Type:  "cup",
		Value: 2,
	}

	converted, err := gram.Convert(toCup)
	if err != nil {
		t.Errorf("Cannot convert unit. Error %s", err)
	}

	if converted.ValueOf() != "2.000000" {
		t.Errorf("Expected 2, got %s", converted.ValueOf())
	}
}
