package interest

import (
	"math"
	"testing"

	"github.com/antonioiubatti93/ophelos/termstructure"
	"github.com/stretchr/testify/assert"
)

func Test_Rate_Discount(t *testing.T) {
	t.Parallel()

	const (
		tm  = 0.5
		tol = 1.0e-15
	)

	for _, tc := range []struct {
		name     string
		rate     Rate
		expected float64
	}{
		{
			"constant term structure/simple discount",
			NewRate(termstructure.NewFlat(0.01), NewSimpleDiscount()),
			1.0 / 1.005,
		},
		{
			"parametric term structure/continuous discount",
			NewRate(termstructure.TermStructureFunc(func(t float64) float64 {
				return math.Log(1.01) / t
			}), NewContinuousDiscount()),
			1.0 / 1.01,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, tc.rate.Discount(tm), tol)
		})
	}
}

func Test_Rate_Shift(t *testing.T) {
	t.Parallel()

	const (
		tm    = 0.5
		value = 0.02
		shift = 0.0001
		tol   = 1.0e-15
	)

	r := NewRate(termstructure.NewFlat(value), NewSimpleDiscount())
	shifted := r.Shift(shift)

	assert.InDelta(t, r.Spot(tm)+shift, shifted.Spot(tm), tol)
}
