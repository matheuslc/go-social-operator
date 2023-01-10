package measurements

import (
	"errors"
)

type Unit float64

type Uniter interface {
	TypeOf() string
	ValueOf() float64
}

var (
	ErrNotConvertible = errors.New("Not convertible")
)

type UnitType struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type Conversor interface {
	Convert(to Uniter) (Uniter, error)
}

func (t UnitType) Convert(to UnitType) (Uniter, error) {
	switch t.Type {
	case "gram":
		g := Gram(t.Value)
		return g.Convert(to)
	case "cup":
		c := Cup(t.Value)
		return c.Convert(to)
	}

	return nil, ErrNotConvertible
}

// ValueOf defines
func (unit UnitType) ValueOf() float64 {
	return unit.Value
}

// ValueOf defines
func (unit UnitType) TypeOf() string {
	return unit.Type
}
