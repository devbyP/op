package handler

import (
	"op/internal/storage"
)

type Service struct {
	GreetingStorage *storage.MessageStorage
}

func New(path string) *Service {
	gstore := storage.NewMessageStorage(path)
	return &Service{
		GreetingStorage: gstore,
	}
}
