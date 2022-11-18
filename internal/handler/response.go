package handler

import "github.com/gin-gonic/gin"

type errorResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status 	string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, msg string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{msg})
}