package interest

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Discount_At(t *testing.T) {
	t.Parallel()

	const (
		r   = 0.02
		tm  = 0.5
		tol = 1.0e-15
	)

	for _, tc := range []struct {
		name     string
		discount Discount
		expected float64
	}{
		{
			"simple",
			NewSimpleDiscount(),
			1 / (1.0 + r*tm),
		},
		{
			"compounded/yearly",
			NewCompoundedDiscount(1),
			math.Pow(1.0+r*tm, -tm),
		},
		{
			"continuous",
			NewContinuousDiscount(),
			math.Exp(-r * tm),
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, tc.discount.At(r, tm), tol)
		})
	}
}
