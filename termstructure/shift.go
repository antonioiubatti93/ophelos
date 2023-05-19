package termstructure

func Shift(ts TermStructure, shift float64) TermStructureFunc {
	return func(t float64) float64 {
		return ts.Value(t) + shift
	}
}
