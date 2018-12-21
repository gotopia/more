package config

import (
	"github.com/spf13/viper"
)

var config = viper.New()

func init() {
	config.SetEnvPrefix("more")
	config.AutomaticEnv()

	config.SetConfigName(Env())
	config.AddConfigPath("config/environments")
	config.ReadInConfig()
}

// Env returns current environment.
func Env() string {
	env := config.GetString("env")
	if env == "" {
		env = "development"
	}
	return env
}

// Development checks whether the server is running in development environment.
func Development() bool {
	return Env() == "development"
}
