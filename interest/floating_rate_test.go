package interest

import (
	"testing"

	"github.com/antonioiubatti93/ophelos/termstructure"
	"github.com/stretchr/testify/assert"
)

func Test_FloatingRate_Spot(t *testing.T) {
	t.Parallel()

	const (
		tm     = 0.5
		value  = 0.02
		spread = 0.01
		tol    = 1.0e-15
	)

	f := NewFloatingRate(termstructure.NewFlat(value), spread)
	assert.InDelta(t, value+spread, f.Spot(tm), tol)
}
