package interest

import (
	"github.com/antonioiubatti93/ophelos/termstructure"
)

type Rate struct {
	ts       termstructure.TermStructure
	discount Discount
}

func NewRate(ts termstructure.TermStructure, discount Discount) Rate {
	return Rate{
		ts:       ts,
		discount: discount,
	}
}

func (r Rate) Spot(t float64) float64 {
	return r.ts.Value(t)
}

func (r Rate) Discount(t float64) float64 {
	return r.discount.At(r.Spot(t), t)
}

func (r Rate) Shift(shift float64) Rate {
	return NewRate(termstructure.Shift(r.ts, shift), r.discount)
}
