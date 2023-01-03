package measurements

type Millimeter Unit

func (ml Millimeter) ValueOf() float64 {
	return float64(ml)
}

func (ml Millimeter) ToCup() Cup {
	return Cup(float64(ml) / 240)
}

func (ml Millimeter) Convert(to Uniter) (Uniter, error) {
	switch to.(type) {
	case Cup:
		return ml.ToCup(), nil
	}

	return nil, ErrNotConvertible
}
