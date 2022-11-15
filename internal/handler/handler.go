package handler

import (
	"net/http"
	"time"

	"github.com/MrDavudov/TestWB/internal/service"
	"github.com/gorilla/mux"
)


type Handler struct {
	http.Server
	router *mux.Router
	services	*service.Service
}

func NewHandler(addr string, services *service.Service) *Handler {
	s := &Handler{
		Server: http.Server{
			Addr:         addr,
			MaxHeaderBytes: 1<<20, // 1MB
            ReadTimeout:  10 * time.Second,
            WriteTimeout: 10 * time.Second,
		},
		router: mux.NewRouter(),
		services:	services,
	}

	s.initRoutes()

	return s
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) initRoutes() {
	h.router.HandleFunc("/", h.handelAllCities)
}

func (h *Handler) handelAllCities(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All cities!\n"))
}