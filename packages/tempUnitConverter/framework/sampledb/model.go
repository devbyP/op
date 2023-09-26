package sampledb

type TempRecord struct {
	ID         int     `json:"id"`
	TempInC    float64 `json:"temp"`
	ReportedAt int64   `json:"reported_at"`
	Location   string  `json:"location"`
}
