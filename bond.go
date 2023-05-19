package main

import "github.com/antonioiubatti93/ophelos/interest"

type Bond struct {
	notional float64
	t        float64
	coupons  []Coupon
}

func NewBond(notional, t float64, coupons ...Coupon) Bond {
	return Bond{
		notional: notional,
		t:        t,
		coupons:  coupons,
	}
}

func (b Bond) Value(rate interest.Rate) float64 {
	npv := b.notional * rate.Discount(b.t)
	for _, c := range b.coupons {
		t := c.At()
		npv += b.notional * c.Value(t) * rate.Discount(t)
	}

	return npv
}
