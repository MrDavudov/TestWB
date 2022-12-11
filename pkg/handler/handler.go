package handler

import (
	"net/http"

	"github.com/MrDavudov/TestWB/pkg/service"
	"github.com/gorilla/mux"
)

type server struct {
	router 		*mux.Router
	services	*service.Service
}

func New(services *service.Service) *server {
	s := &server{
		router:  	mux.NewRouter(),
		services:	services,
	}

	s.initRoutes()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) initRoutes() {
	s.router.HandleFunc("/", s.home).Methods("GET")
	s.router.HandleFunc("/all", s.allCities).Methods("GET")
	s.router.HandleFunc("/all/{city}", s.getCity).Methods("GET")
	s.router.HandleFunc("/city", s.cityCreate()).Methods("POST")
	s.router.HandleFunc("/city", s.cityDelete()).Methods("DELETE")
}