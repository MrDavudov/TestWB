package service

import (
	"fmt"
	"testing"

	"github.com/MrDavudov/TestWB/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestGetCityAPI(t *testing.T) {
	// Настройка тестовых данных
	tTable := []struct {
		city     string
		expected model.Weather
	}{
		{
			city: "Москва",
			expected: model.Weather{
				Name:    "Moscow",
				Lat:     55.7504461,
				Lon:     37.6174943,
				Country: "RU",
			},
		},
		{
			city: "Moscow",
			expected: model.Weather{
				Name:    "Moscow",
				Lat:     55.7504461,
				Lon:     37.6174943,
				Country: "RU",
			},
		},
		{
			city:     "",
			expected: model.Weather{},
		},
		{
			city:     "dfsdfsdf",
			expected: model.Weather{},
		},
	}
	// Вызов тестированого кода
	for _, tCase := range tTable {
		result, _ := GetCityAPI(tCase.city)

		// t.Logf("Calling GetCity(%s), resilt %v\n", tCase.city, result)
		// Проверка
		t.Run(tCase.city, func(t *testing.T) {
			assert.Equal(t, tCase.expected, result, fmt.Sprintf("Incorrect ressult, Expect %v, got %v",
				tCase.expected, result))
		})
	}
}

func TestGetDataTempAllAPI(t *testing.T) {
	// Настройка тестовых данных
	tTable := []struct {
		in     []model.Weather
		expected []model.Weather
	}{
		{
			in: []model.Weather{
				{
					Name:    "Moscow",
					Lat:     55.7504461,
					Lon:     37.6174943,
					Country: "RU",
					DtTemp:  nil,
				},
			},
			// expected может быть неактуальный, так как погода все время меняется
			expected: []model.Weather{
				{
					Name:    "Moscow",
					Lat:     55.7504461,
					Lon:     37.6174943,
					Country: "RU",
					DtTemp: []model.DtTemp{
						{
							Dt:   "2022-12-12",
							Temp: 1.89,
						},
						{
							Dt:   "2022-12-13",
							Temp: 0.18,
						},
						{
							Dt:   "2022-12-14",
							Temp: -3.34,
						},
						{
							Dt:   "2022-12-15",
							Temp: -7.58,
						},
						{
							Dt:   "2022-12-16",
							Temp: -4.46,
						},
					},
				},
			},
		},
		{
			in: []model.Weather{},
			expected: []model.Weather{},
		},
		{
			in: []model.Weather{
				{
					Name: "Moscow",
				},
			},
			expected: nil,
		},
	}
	// Вызов тестированого кода
	for _, tCase := range tTable {
		result := GetDataTempAllAPI(tCase.in)

		// t.Logf("Calling GetCity(%v), resilt %v\n", tCase.city, result)
		// Проверка
		t.Run("Test", func(t *testing.T) {
			assert.Equal(t, tCase.expected, result, fmt.Sprintf("Incorrect ressult, Expect %v, got %v",
				tCase.expected, result))
		})
	}
}

func TestGetDataTempCityAPI(t *testing.T) {
	// Настройка тестовых данных
	tTable := []struct {
		in     model.Weather
		expected model.Weather
	}{
		{
			in: model.Weather{
				Name:    "Moscow",
				Lat:     55.7504461,
				Lon:     37.6174943,
				Country: "RU",
				DtTemp: []model.DtTemp{},
			},
			// expected может быть неактуальный, так как погода все время меняется
			expected: model.Weather{
				Name:    "Moscow",
				Lat:     55.7504461,
				Lon:     37.6174943,
				Country: "RU",
				DtTemp: []model.DtTemp{
					{
						Dt:   "2022-12-12",
						Temp: 1.89,
					},
					{
						Dt:   "2022-12-13",
						Temp: 0.18,
					},
					{
						Dt:   "2022-12-14",
						Temp: -3.34,
					},
					{
						Dt:   "2022-12-15",
						Temp: -7.58,
					},
					{
						Dt:   "2022-12-16",
						Temp: -4.46,
					},
				},
			},
		},
		{
			in:     model.Weather{},
			expected: model.Weather{},
		},
		{
			in: model.Weather{
				Name: "Moscow",
			},
			expected: model.Weather{},
		},
	}
	// Вызов тестированого кода
	for _, tCase := range tTable {
		result, _ := GetDataTempCityAPI(tCase.in)

		// t.Logf("Calling GetCity(%s), resilt %v\n", tCase.city, result)
		// Проверка
		t.Run(tCase.in.Name, func(t *testing.T) {
			assert.Equal(t, tCase.expected, result, fmt.Sprintf("Incorrect ressult, Expect %v, got %v",
				tCase.expected, result))
		})
	}
}
