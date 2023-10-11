package sampledb

import (
	"errors"

	"tempUnitConverter/domain"
)

func (s *Storage) GetTemperatureReport(id int) (*domain.TemperatureReport, error) {
	if id >= len(s.DB.Data) || id < 0 {
		return nil, errors.New("report not found")
	}
	val := s.DB.Data[id]
	return val.toDomainTempReport(), nil
}

func (s *Storage) GetAllTemperatureReports() ([]*domain.TemperatureReport, error) {
	rep := make([]*domain.TemperatureReport, len(s.DB.Data))
	for _, val := range s.DB.Data {
		rep = append(rep, val.toDomainTempReport())
	}
	return rep, nil
}
