package app

import "name_message/core/domain"

// secondary port
type DbPort interface {
	GetName() (*domain.NameMessage, error)
	SaveName(name string) (int, error)
}

// primary port
type NamePorrt interface {
	GetName() (*GetNameResponse, error)
	SaveName(req *SavaNameRequest) (*SaveNameResponse, error)
}
