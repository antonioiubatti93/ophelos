package interest

import "math"

type Discount interface {
	At(r, t float64) float64
}

type SimpleDiscount struct{}

var _ Discount = SimpleDiscount{}

func NewSimpleDiscount() SimpleDiscount {
	return SimpleDiscount{}
}

func (d SimpleDiscount) At(r, t float64) float64 {
	return 1.0 / (1.0 + r*t)
}

type CompoundedDiscount struct {
	frequency float64
}

var _ Discount = CompoundedDiscount{}

func NewCompoundedDiscount(f int) CompoundedDiscount {
	return CompoundedDiscount{
		frequency: float64(f),
	}
}

func (d CompoundedDiscount) At(r, t float64) float64 {
	return math.Pow(1.0+r*t/d.frequency, -d.frequency*t)
}

type ContinuousDiscount struct{}

func NewContinuousDiscount() ContinuousDiscount {
	return ContinuousDiscount{}
}

var _ Discount = ContinuousDiscount{}

func (d ContinuousDiscount) At(r, t float64) float64 {
	return math.Exp(-r * t)
}
