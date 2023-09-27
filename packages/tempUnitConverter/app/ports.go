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
	FtoC(f domain.F) (domain.C, error)
	CtoF(c domain.C) (domain.F, error)
	GetReport(id int) (*domain.TemperatureReport, error)
	GetReports() ([]*domain.TemperatureReport, error)
}
