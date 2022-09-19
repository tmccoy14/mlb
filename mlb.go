package main

import (
	"github.com/tmccoy14/mlb/internal/config"
	"github.com/tmccoy14/mlb/router"
)

func main() {
	// Echo instance
	e := router.New()

	// Load application environment variables
	config, err := config.LoadConfig(e)
	if err != nil {
		e.Logger.Fatal("Failed to load config:", err)
	}

	// Start server
	e.Logger.Fatal(e.Start(config.Port))
}
