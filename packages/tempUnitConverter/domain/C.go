package domain

type C float64

func (c C) Float64() float64 {
	return float64(c)
}

func (c C) Unit() string {
	return "C"
}

func (c C) ToF() F {
	return F((c * 9 / 5) + 32)
}
