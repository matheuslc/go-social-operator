package measurements

import "fmt"

// Gram defines the gram unit type
// e.g 1.0
type Gram Unit

func (gram Gram) ValueOf() string {
	return fmt.Sprintf("%f", float64(gram))
}

func (gram Gram) ToCup() Cup {
	return Cup(float64(gram) / 240)
}

func (gram Gram) Convert(to UnitType) (Uniter, error) {
	switch to.Type {
	case "cup":
		return gram.ToCup(), nil
	}

	return nil, ErrNotConvertible
}
