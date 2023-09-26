package handler

import (
	"context"

	"op/greetingapi"
)

type GreetingService struct {
	service *Service
}

func NewGreetingService(service *Service) *GreetingService {
	return &GreetingService{
		service: service,
	}
}
