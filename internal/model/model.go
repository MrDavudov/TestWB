package model

type Weather struct {
	Name    string
	Lat     float64
	Lon     float64
	Country string
	DtTemp	[]DtTemp
}

type DtTemp struct {
	Dt   string
	Temp float64
}