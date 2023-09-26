package domain

type C float64

func (c C) Temp() float64 {
	return float64(c)
}

func (c C) ToF() F {
	return F((c * 9 / 5) + 32)
}
