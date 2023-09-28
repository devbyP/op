package app

import (
	"tempUnitConverter/domain"
)

// driven port (db)
type ITemperatureStorage interface {
	SaveTemperatureReport(temp domain.C) (int, error)
	GetAllTemperatureReports() ([]*domain.TemperatureReport, error)
	GetTemperatureReport(id int) (*domain.TemperatureReport, error)
}

// driving port
type AppPort interface {
	SaveFtoC(f domain.F) (int, error)
	SaveC(c domain.C) (int, error)
	FToC(f domain.F) domain.C
	CToF(c domain.C) domain.F
	GetReportByID(id int) (*domain.TemperatureReport, error)
	GetAllReports() ([]*domain.TemperatureReport, error)
}
