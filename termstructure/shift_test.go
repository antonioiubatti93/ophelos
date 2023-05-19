package termstructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Shift(t *testing.T) {
	t.Parallel()

	const (
		tm    = 0.5
		value = 0.01
		shift = 0.0001
		tol   = 1.0e-15
	)

	ts := NewFlat(value)
	assert.InDelta(t, value+shift, Shift(ts, shift).Value(tm), tol)
}
