package app

import (
	"tempUnitConverter/domain"
)

func (t *Temperature) GetReportByID(id int) (*domain.TemperatureReport, error) {
	return t.db.GetTemperatureReport(id)
}

func (t *Temperature) GetAllReports() ([]*domain.TemperatureReport, error) {
	return t.db.GetAllTemperatureReports()
}
