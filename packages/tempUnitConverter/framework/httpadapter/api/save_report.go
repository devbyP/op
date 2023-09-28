package api

type RequestSaveReport struct {
	Temp float64 `json:"temp"`
	Unit string  `json:"unit"`
}

type ResponseSaveReport struct {
	ID int `json:"id"`
}
