package api

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/tmccoy14/mlb/handlers"
	"github.com/tmccoy14/mlb/validator"
)

// Template data
type Template struct {
	templates *template.Template
}

// Render HTML template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Main route group for landing page and healthz
func MainGroup(e *echo.Echo) {

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.GET("/", handlers.Landing)
	e.GET("/healthz", handlers.HealthCheck)
}

// Lookup route group to search players, teams, and stats
func LookupGroup(g *echo.Group, e *echo.Echo) {

	validator.InitValidator(e)

	g.GET("/players", handlers.Players)
	g.GET("/teams", handlers.Teams)
	g.GET("/stats", handlers.Stats)
}
