package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
	Status  int `json:"status"`
}

func HealthCheck(c echo.Context) error {

	// Set resp to the go struct pointer
	resp := HealthCheckResponse{
		Message: "OK!",
		Status: http.StatusOK,
	}

	// Return health check message and status code
	return c.JSON(http.StatusOK, resp)
}
