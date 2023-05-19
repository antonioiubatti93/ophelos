package termstructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testTermStructure struct {
	a float64
}

var _ TermStructure = testTermStructure{}

func (ts testTermStructure) Value(t float64) float64 {
	return ts.a * t
}

func Test_TermStructure_Value(t *testing.T) {
	t.Parallel()

	const (
		tm  = 1.0
		tol = 1.0e-15
	)

	for _, tc := range []struct {
		name     string
		ts       TermStructure
		expected float64
	}{
		{
			"flat",
			NewFlat(0.02),
			0.02,
		},
		{
			"linear",
			testTermStructure{a: 0.01},
			0.01 * tm,
		},
	} {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			assert.InDelta(t, tc.expected, tc.ts.Value(tm), tol)
		})
	}
}
