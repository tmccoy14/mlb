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
	e.Use(middleware.CORS())

	// Add static files and assets
	e.File("/favicon.ico", "assets/favicon.ico")
	e.Static("/static", "static")

	// Create route groups
	lookupGroup := e.Group("/lookup")

	// Set main routes
	api.MainGroup(e)

	// Set group routes
	api.LookupGroup(lookupGroup, e)

	return e
}
