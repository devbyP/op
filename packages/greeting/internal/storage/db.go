package storage

import (
	"csv_file_database/csvdb"
)

func GetDB[T any](path string) (*csvdb.Database[T], error) {
	return csvdb.OpenCsvDatabase[T](path)
}

func GenerateIncrementalID[T any](data []T) int {
	return len(data)
}
