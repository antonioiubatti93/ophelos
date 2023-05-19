package main

import "github.com/antonioiubatti93/ophelos/interest"

type Coupon interface {
	Value(t float64) float64
	At() float64
}

type FixedCoupon struct {
	rate float64
	t    float64
}

var _ Coupon = FixedCoupon{}

func NewFixedCoupon(rate, t float64) FixedCoupon {
	return FixedCoupon{
		rate: rate,
		t:    t,
	}
}

func (c FixedCoupon) Value(_ float64) float64 {
	return c.rate
}

func (c FixedCoupon) At() float64 {
	return c.t
}

type FloatingCoupon struct {
	rate interest.FloatingRate
	t    float64
}

func NewFloatingCoupon(rate interest.FloatingRate, t float64) FloatingCoupon {
	return FloatingCoupon{
		rate: rate,
		t:    t,
	}
}

var _ Coupon = FloatingCoupon{}

func (c FloatingCoupon) Value(t float64) float64 {
	return c.rate.Spot(t)
}

func (c FloatingCoupon) At() float64 {
	return c.t
}
