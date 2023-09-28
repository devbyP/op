package app

type Temperature struct {
	db ITemperatureStorage
}

func NewTemperature(db ITemperatureStorage) *Temperature {
	return &Temperature{db: db}
}
