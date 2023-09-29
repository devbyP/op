package adapter

import (
	"op/internal/handler"
	"op/internal/handler/ports"
)

type Adapter struct {
	greetingService ports.IGreeting
}

func New(service *handler.Service) *Adapter {
	return &Adapter{
		greetingService: handler.NewGreetingService(service),
	}
}
