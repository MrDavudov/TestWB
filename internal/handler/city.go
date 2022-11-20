package handler

import (
	"net/http"

	"github.com/MrDavudov/TestWB/internal/model"
	"github.com/gin-gonic/gin"
)

func (s *handler) home(c *gin.Context) {
	c.JSON(http.StatusOK, StatusResponse{
		Status: "Start!",
	})
}

// Get all city
func (s *handler) allCities(c *gin.Context) {
	citiesAll, err := s.services.GetAllCities()
	if err != nil  {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, citiesAll)
}

// Get city weather
func (s *handler) getCity(c *gin.Context) {
	var request model.Weather

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	city, err := s.services.GetCity(request.Name)
	if err != nil  {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, city)
}

// Create city
func (s *handler) cityCreate(c *gin.Context) {
	var request model.Weather

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	city, err := s.services.Save(request.Name); 
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, city)
}

// Delete city
func (s *handler) cityDelete(c *gin.Context) {
	var request model.Weather

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := s.services.City.Delete(request.Name)
	if err != nil  {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "Delete: " + request.Name,
	})
}