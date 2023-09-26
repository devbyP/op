package sampledb

import (
	"csv_file_database/csvdb"
	"tempUnitConverter/app"
)

var _ app.ITemperatureStorage = &Storage{}

type Storage struct {
	db *csvdb.Database[*TempRecord]
}

func New() *Storage {
	db, err := csvdb.OpenCsvDatabase[*TempRecord]("./temp.csv")
	if err != nil {
		panic(err)
	}
	return &Storage{db: db}
}
