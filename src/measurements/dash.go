package measurements

type Dash Unit

func (dash Dash) ValueOf() float64 {
	return float64(dash)
}

func (dash Dash) TypeOf() string {
	return "dash"
}

func (dash Dash) ToCup() Cup {
	// 1 cup = 240 / 5 = 48 grams
	return Cup(float64(dash) / 48)
}

func (dash Dash) ToGram() Gram {
	return Gram(dash * 5)
}

func (dash Dash) Convert(to UnitType) (Uniter, error) {
	switch to.Type {
	case "cup":
		return dash.ToCup(), nil
	case "gram":
		return dash.ToGram(), nil
	}

	return nil, ErrNotConvertible
}
