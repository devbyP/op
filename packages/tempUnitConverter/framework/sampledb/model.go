package sampledb

import (
	"time"

	"tempUnitConverter/domain"
)

type TempRecord struct {
	ID         int     `json:"id"`
	TempInC    float64 `json:"temp"`
	ReportedAt int64   `json:"reported_at"`
	Location   string  `json:"location"`
}

func (t TempRecord) toDomainReport() *domain.TemperatureReport {
	return &domain.TemperatureReport{
		ID:         t.ID,
		TempInC:    t.TempInC,
		ReportedAt: time.Unix(t.ReportedAt, 0),
		Location:   t.Location,
	}
}
