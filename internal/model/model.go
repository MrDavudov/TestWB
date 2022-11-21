package model

type Weather struct {
	Name    string   `json:"name" db:"city" binding:"required"`
	Lat     float64  `json:"lat"`
	Lon     float64  `json:"lon"`
	Country string   `json:"country"`
	DtTemp  []DtTemp `json:"dt_temp"`
}

type DtTemp struct {
	Dt   string  `json:"dt" db:"dt"`
	Temp float64 `json:"temp" db:"temp"`
}
