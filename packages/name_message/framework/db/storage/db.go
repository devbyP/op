package storage

import (
	"csv_file_database/csvdb"
	"name_message/framework/db/model"
)

type (
	Storage struct {
		db *csvdb.Database[*model.NameMessageDB]
	}
	NameStorage struct {
		store *Storage
	}
)

func New(path string) *Storage {
	db, err := csvdb.OpenCsvDatabase[*model.NameMessageDB](path)
	if err != nil {
		panic(err)
	}
	return &Storage{db: db}
}
