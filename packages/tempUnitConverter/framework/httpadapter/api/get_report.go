package api

import (
	"tempUnitConverter/domain"
)

type ReponseGetReport struct {
	Data *domain.TemperatureReport `json:"data"`
}

type ResponseGetAllReports struct {
	Data []*domain.TemperatureReport `json:"data"`
}
