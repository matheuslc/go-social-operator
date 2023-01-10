package measurements

// Gram defines the gram unit type
// e.g 1.0
type Gram Unit

func (gram Gram) ValueOf() float64 {
	return float64(gram)
}

func (gram Gram) TypeOf() string {
	return "gram"
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
