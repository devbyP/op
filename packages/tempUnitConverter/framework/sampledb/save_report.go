package sampledb

import (
	"tempUnitConverter/domain"
	"time"
)

func (s *Storage) SaveTemperatureReport(c domain.C) (int, error) {
	curId := len(s.db.Data)
	tr := &TempRecord{
		ID:         curId,
		TempInC:    c.Float64(),
		ReportedAt: time.Now().Unix(),
		Location:   "Bangkok",
	}
	s.db.Data = append(s.db.Data, tr)
	return curId, nil
}
