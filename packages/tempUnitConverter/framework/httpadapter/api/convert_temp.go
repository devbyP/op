package api

type RequestConverTemp struct {
	Unit string  `json:"unit"`
	Temp float64 `json:"temp"`
}

type ReponseCovertTemp struct {
	Unit string  `json:"unit"`
	Temp float64 `json:"temp"`
}
