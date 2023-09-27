package domain

type TemperatureUnit interface {
	Float64() float64
	Unit() string
}
