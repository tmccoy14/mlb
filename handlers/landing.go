package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func Landing(c echo.Context) error {
	
	// Return rendered HTML template
	return c.Render(http.StatusOK, "landing.html", nil)
}
