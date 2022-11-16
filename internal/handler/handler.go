package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MrDavudov/TestWB/internal/model"
	"github.com/MrDavudov/TestWB/internal/service"
	"github.com/gorilla/mux"
)

type server struct {
	http.Server
	router *mux.Router
	services	*service.Service
}

type statusResponse struct {
	Msg	string `json:"msg" binding:"required"`
}

func NewHandler(port string, services *service.Service) *server {
	s := &server{
		Server: http.Server{
			Addr:         	port,
			MaxHeaderBytes: 1<<20, // 1MB
            ReadTimeout:  	10 * time.Second,
            WriteTimeout: 	10 * time.Second,
		},
		router: mux.NewRouter(),
		services:	services,
	}

	s.initRoutes()

	//установить обработчик http-сервера
	s.Handler = s.router

	return s
}

func (s *server) initRoutes() {
	s.router.HandleFunc("/", s.handelHome()).Methods("GET")
	s.router.HandleFunc("/all", s.handelAllCities()).Methods("GET")
	s.router.HandleFunc("/city", s.handleCityCreate()).Methods("POST")
	s.router.HandleFunc("/city", s.handleCityDelete()).Methods("DELETE")
}

func (s *server) handelHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request)  {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Start!\n"))
	}
}

func (s *server) handelAllCities() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request)  {
		citiesAll, err := s.services.DataTemp.GetAllCities()
		if err != nil  {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		rawDataOut, err := json.Marshal(&citiesAll)
		if err != nil {
			return
		}

		w.Write(rawDataOut)

		// s.respond(w, r, http.StatusCreated, w.Write(rawDataOut))
	}
}

func (s *server) handleCityCreate() http.HandlerFunc {
	type request struct {
		Name	string	`json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		city := model.Weather{
			Name: req.Name,
		}
		if err := s.services.City.Save(city.Name) ; err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		city = service.GetCity(city.Name)
		rawDataOut, err := json.Marshal(&city)
		if err != nil {
			return
		}
		s.respond(w, r, http.StatusCreated, rawDataOut)
	}
}

func (s *server) handleCityDelete() http.HandlerFunc {
	type request struct {
		Name string	`json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		city := req.Name

		if err := s.services.City.Delete(city) ; err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, w)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w)
	}
}