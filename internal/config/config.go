package config

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Config struct {
	ProjectId  string `mapstructure:"PROJECT_ID"`
	Collection string `mapstructure:"COLLECTION"`
	Port       string `mapstructure:"PORT"`
}

func LoadConfig(e *echo.Echo) (config Config, err error) {

	// Load environment variables config
	viper.AddConfigPath(".config")
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	// Read in environment variables from config
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	// Iterate over fields to ensure environment variables are set
	c := Config{
		ProjectId:  config.ProjectId,
		Collection: config.Collection,
		Port:       config.Port,
	}
	verify := reflect.ValueOf(c)
	for i := 0; i < verify.NumField(); i++ {
		key := verify.Type().Field(i).Name
		value := verify.Field(i)

		if value.IsZero() {
			e.Logger.Fatalf("%s environment variable is empty.", key)
		}
	}

	return config, err
}
