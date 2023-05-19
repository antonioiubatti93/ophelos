package model

import (
	"fmt"

	"github.com/antonioiubatti93/ophelos/termstructure"
	interp "github.com/edgelaboratories/interpolator"
	"golang.org/x/exp/slices"
)

type interpolator interface {
	Value(x float64) float64
}

type Discrete struct {
	interpolator
}

var _ termstructure.TermStructure = Discrete{}

func NewDiscrete(tvs ...TenorValue) (Discrete, error) {
	xys := make(interp.XYs, 0, len(tvs))
	for _, tv := range tvs {
		xys = append(xys, interp.XY{
			X: tv.tenor.Value(),
			Y: tv.value,
		})
	}

	slices.SortFunc(xys, func(a, b interp.XY) bool {
		return a.X < b.X
	})

	m, err := interp.NewPiecewiseLinearThreshold(xys)
	if err != nil {
		return Discrete{}, fmt.Errorf("could not build interpolator: %w", err)
	}

	return Discrete{m}, nil
}
