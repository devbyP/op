package sampledb

import (
	"errors"
	"tempUnitConverter/domain"
)

func (s *Storage) GetTemperatureReport(id int) (*domain.TemperatureReport, error) {
	d := s.DB.Data
	if id > len(d) {
		return nil, errors.New("report not found")
	}
	val := d[id]
	return val.toDomainReport(), nil
}

func (s *Storage) GetAllTemperatureReports() ([]*domain.TemperatureReport, error) {
	d := make([]*domain.TemperatureReport, len(s.DB.Data))
	for i, val := range s.DB.Data {
		d[i] = val.toDomainReport()
	}
	return d, nil
}
