package interest

import "github.com/antonioiubatti93/ophelos/termstructure"

type FloatingRate struct {
	ts     termstructure.TermStructure
	spread float64
}

func NewFloatingRate(ts termstructure.TermStructure, spread float64) FloatingRate {
	return FloatingRate{
		ts:     ts,
		spread: spread,
	}
}

func (f FloatingRate) Spot(t float64) float64 {
	return f.ts.Value(t) + f.spread
}
