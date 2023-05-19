package model

type Tenor int

const (
	ON Tenor = iota
	W1
	M1
	M3
	M6
	M12
)

func (t Tenor) Value() float64 {
	return [...]float64{
		1.0 / 360.0,
		1.0 / 52.0,
		1.0 / 12.0,
		0.25,
		0.5,
		1.0,
	}[t]
}

type TenorValue struct {
	tenor Tenor
	value float64
}

func NewTenorValue(tenor Tenor, value float64) TenorValue {
	return TenorValue{
		tenor: tenor,
		value: value,
	}
}
