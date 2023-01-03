package measurements

// UnitToGram creates a new gram with an defined alue
type Cup Gram

func (cup Cup) ValueOf() float64 {
	return float64(cup)
}

// Convert Cup to Gram
func (cup Cup) ToGram() Gram {
	return Gram(int(cup) * 240)
}

func (cup Cup) Convert(to Uniter) (Uniter, error) {
	switch to.(type) {
	case Gram:
		return cup.ToGram(), nil
	}

	return nil, ErrNotConvertible
}
