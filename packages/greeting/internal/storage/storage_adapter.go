package storage

import (
	"csv_file_database/csvdb"
	"op/internal/model"
)

type (
	Storage[T ModelType] struct {
		db *csvdb.Database[T]
	}

	ModelType interface {
		*model.MessageData
	}
)

func New[T ModelType](path string) (*Storage[T], error) {
	db, err := GetDB[T](path)
	if err != nil {
		return nil, err
	}
	return &Storage[T]{
		db: db,
	}, nil
}

func Must[T ModelType](s *Storage[T], err error) *Storage[T] {
	if err != nil {
		panic(err)
	}
	return s
}
