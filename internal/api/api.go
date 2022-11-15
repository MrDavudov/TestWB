package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/MrDavudov/TestWB/internal/model"
	"go.uber.org/zap"
)

const base = "http://api.openweathermap.org"
const pathCity = "/geo/1.0/direct?"
const pathWeather = "/data/2.5/forecast?"
const apiKeys = "&appid=90f2edc318c106c65581f4052ad16c6f"

type City struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

func GetCity(city string) *model.Weather {
	var logger, _ = zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	resp, err := http.Get(base + pathCity + "q=" + city + apiKeys)
	if err != nil {
		sugar.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		sugar.Fatal(err)
	}
	if string(body) == "[]" {
		sugar.Fatal("error no such city")
	}

	obj := []City{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		sugar.Fatal(err)
	}

	Weather := model.Weather{
		Name:		obj[0].Name,
		Lat:		obj[0].Lat,
		Lon:		obj[0].Lon,
		Country:	obj[0].Country,
	}

	return &Weather
}

type DataTemp struct {
	List []struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Data string `json:"dt_txt"`
	} `json:"list"`
}

func GetDataTemp(weather model.Weather) {
	var logger, _ = zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	lat := fmt.Sprintf("lat=%f", weather.Lat)
	lon := fmt.Sprintf("&lon=%f", weather.Lon)

	resp, err := http.Get(base + pathWeather + lat + lon + "&units=metric" + apiKeys)
	if err != nil {
		sugar.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		sugar.Fatal(err)
	}

	obj := DataTemp{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		sugar.Fatal(err)
	}

	for i := range obj.List {
		if strings.Contains(obj.List[i].Data, "12:00") {
			d := model.DtTemp {
				Dt: obj.List[i].Data,
				Temp: obj.List[i].Main.Temp,
			}
			weather.DtTemp = append(weather.DtTemp, d)
		}
	}
}