package sampledb

import (
	"csv_file_database/csvdb"
	"tempUnitConverter/app"
)

var _ app.ITemperatureStorage = &Storage{}

type Storage struct {
	DB *csvdb.Database[*TempRecord]
}

func New(path string) *Storage {
	db, err := csvdb.OpenCsvDatabase[*TempRecord](path)
	if err != nil {
		panic(err)
	}
	return &Storage{DB: db}
}
