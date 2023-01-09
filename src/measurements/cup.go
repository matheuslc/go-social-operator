package measurements

import "fmt"

// UnitToGram creates a new gram with an defined alue
type Cup Gram

func (cup Cup) ValueOf() string {
	return fmt.Sprintf("%f", float64(cup))
}

func (cup Cup) String() string {
	return fmt.Sprintf("%f", cup)
}

// Convert Cup to Gram
func (cup Cup) ToGram() Gram {
	return Gram(int(cup) * 240)
}

func (cup Cup) Convert(to UnitType) (Uniter, error) {
	switch to.Type {
	case "gram":
		return cup.ToGram(), nil
	}

	return nil, ErrNotConvertible
}
