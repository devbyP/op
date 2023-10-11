package sampledb

import (
	"time"

	"tempUnitConverter/domain"
)

func (s *Storage) SaveTemperatureReport(c domain.C) (int, error) {
	currId := len(s.DB.Data)
	val := &TempRecord{
		ID:         currId,
		TempInC:    c.Float64(),
		Location:   "Bangkok",
		ReportedAt: time.Now().Unix(),
	}
	s.DB.Data = append(s.DB.Data, val)
	return currId, nil
}
