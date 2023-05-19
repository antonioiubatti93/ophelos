package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tenor_Value(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1.0/360.0, ON.Value())
	assert.Equal(t, 1.0/52.0, W1.Value())
	assert.Equal(t, 1.0/12.0, M1.Value())
	assert.Equal(t, 0.25, M3.Value())
	assert.Equal(t, 0.5, M6.Value())
	assert.Equal(t, 1.0, M12.Value())
}

func Test_NewTenorValue(t *testing.T) {
	t.Parallel()

	const (
		tenor = M1
		value = 0.01
	)

	tv := NewTenorValue(tenor, value)
	assert.Equal(t, tenor, tv.tenor)
	assert.Equal(t, value, tv.value)
}
