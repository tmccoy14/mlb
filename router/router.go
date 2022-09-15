package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	api "github.com/tmccoy14/mlb/api"
)

func New() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Add static files and assets
	e.File("/favicon.ico", "static/img/favicon.ico")
	e.Static("/static", "static")

	// Create route groups
	lookupGroup := e.Group("/lookup")

	// Set main routes
	api.MainGroup(e)

	// Set group routes
	api.LookupGroup(lookupGroup, e)

	return e
}
