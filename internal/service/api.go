package service

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

func GetCity(city string) (model.Weather, error) {
	type City struct {
		Name    string  `json:"name"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
		Country string  `json:"country"`
	}

	resp, err := http.Get(base + pathCity + "q=" + city + apiKeys)
	if err != nil {
		return model.Weather{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Weather{}, err
	}
	if string(body) == "[]" {
		return model.Weather{}, fmt.Errorf("error no such city")
	}

	obj := []City{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		return model.Weather{}, err
	}

	return model.Weather{
		Name:		obj[0].Name,
		Lat:		obj[0].Lat,
		Lon:		obj[0].Lon,
		Country:	obj[0].Country,
	}, nil
}

func GetDataTemp(w []model.Weather) []model.Weather {
	type DataTemp struct {
		List []struct {
			Main struct {
				Temp float64 `json:"temp"`
			} `json:"main"`
			Data string `json:"dt_txt"`
		} `json:"list"`
	}

	var logger, _ = zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	for i := range w {
		lat := fmt.Sprintf("lat=%f", w[i].Lat)
		lon := fmt.Sprintf("&lon=%f", w[i].Lon)

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

		for j := range obj.List {
			if strings.Contains(obj.List[j].Data, "12:00") {
				d := model.DtTemp {
					Dt: obj.List[j].Data,
					Temp: obj.List[j].Main.Temp,
				}
				w[i].DtTemp = append(w[i].DtTemp, d)
			}
		}
	}

	return w
}