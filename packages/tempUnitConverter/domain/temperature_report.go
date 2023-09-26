package domain

import "time"

type TemperatureReport struct {
	ID         int       `json:"id"`
	TempInC    float64   `json:"temp"`
	ReportedAt time.Time `json:"reported_at"`
	Location   string    `json:"location"`
}
