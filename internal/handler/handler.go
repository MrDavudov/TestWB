package handler

import (
	"net/http"
	"time"

	"github.com/MrDavudov/TestWB/internal/service"
	"github.com/gorilla/mux"
)


type server struct {
	http.Server
	router *mux.Router
	services	*service.Service
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
	s.router.HandleFunc("/", s.handelAllCities)
}

func (s *server) handelAllCities(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All cities!\n"))
}