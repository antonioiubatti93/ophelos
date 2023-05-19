package termstructure

type TermStructure interface {
	Value(t float64) float64
}

type TermStructureFunc func(t float64) float64

var _ TermStructure = TermStructureFunc(nil)

func (f TermStructureFunc) Value(t float64) float64 {
	return f(t)
}

func NewFlat(r float64) TermStructureFunc {
	return func(_ float64) float64 {
		return r
	}
}
