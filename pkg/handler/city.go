package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) home(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "Start")
}

// Get all city
func (s *server) allCities(w http.ResponseWriter, r *http.Request) {
	citiesAll, err := s.services.GetAllCities()
	if err != nil  {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, citiesAll)
}

// Get city weather
func (s *server) getCity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    req := vars["city"]

	city, err := s.services.GetCity(req)
	if err != nil  {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, city)
}

// Create city
func (s *server) cityCreate() http.HandlerFunc {
	type request struct {
		City	string	`json:"city"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		city, err := s.services.Save(req.City)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, city)
	}
}

// Delete city
func (s *server) cityDelete() http.HandlerFunc {
	type request struct {
		City	string	`json:"city"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			newErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err := s.services.City.Delete(req.City)
		if err != nil {
			newErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, "Delete: " + req.City)
	}
}