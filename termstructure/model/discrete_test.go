package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewDiscrete_NoPoints(t *testing.T) {
	t.Parallel()

	_, err := NewDiscrete()
	assert.Error(t, err)
}

func Test_NewDiscrete(t *testing.T) {
	t.Parallel()

	ts, err := NewDiscrete(
		NewTenorValue(ON, 0.01),
		NewTenorValue(M1, 0.02),
	)
	require.NoError(t, err)
	assert.InDelta(t, 0.02, ts.Value(M1.Value()), 1.0e-15)
}
