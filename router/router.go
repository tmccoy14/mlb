package router

import (
	api "github.com/tmccoy14/mlb/api"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Add favicon
	e.File("/favicon.ico", "assets/favicon.ico")

	// Create route groups
	lookupGroup := e.Group("/lookup")

	// Set main routes
	api.MainGroup(e)

	// Set group routes
	api.LookupGroup(lookupGroup, e)

	return e
}
