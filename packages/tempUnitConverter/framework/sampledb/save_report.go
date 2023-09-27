package sampledb

import (
	"tempUnitConverter/domain"
	"time"
)

func (s *Storage) SaveTemperatureReport(c domain.C) (int, error) {
	curId := len(s.DB.Data)
	tr := &TempRecord{
		ID:         curId,
		TempInC:    c.Float64(),
		ReportedAt: time.Now().Unix(),
		Location:   "Bangkok",
	}
	s.DB.Data = append(s.DB.Data, tr)
	return curId, nil
}
