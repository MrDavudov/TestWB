package handler

import (
	"net/http"

	"github.com/MrDavudov/TestWB/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	httpServer 	*http.Server
	router 		*gin.Engine
	services	*service.Service
}

func NewHandler(port string, services *service.Service) *handler {
	s := &handler{
		router: gin.Default(),
		services:	services,
	}

	s.initRoutes()

	return s
}

func (s *handler) initRoutes() *gin.Engine {
	api := s.router.GET("/", s.home)
	{
		api.GET("/all", s.allCities)
		api.GET("all/:city", s.getCity)
		api.POST("/city", s.cityCreate)
		api.DELETE("/city", s.cityDelete)
	}
	
	return s.router
}
