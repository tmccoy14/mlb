package main

import (
	"github.com/tmccoy14/mlb/router"
)

func main() {
	// Echo instance
	e := router.New()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
