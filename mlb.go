package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tmccoy14/mlb/router"
)

func main() {
	// Echo instance
	e := router.New()

	// Configure environment variables from configuration file
	viper.SetConfigName("configghgh.yaml")
	viper.AddConfigPath(".config")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		e.Logger.Error("hjh")
		fmt.Printf("Error reading config file, %s", err)
	}

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
