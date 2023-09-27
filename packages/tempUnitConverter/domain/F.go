package domain

type F float64

func (f F) Float64() float64 {
	return float64(f)
}

func (f F) Unit() string {
	return "F"
}

func (f F) ToC() C {
	return C((f - 32) * 5 / 9)
}
