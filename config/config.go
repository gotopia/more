package config

import (
	"strings"

	"github.com/spf13/viper"
)

var config = viper.New()

func init() {
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
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

// QA checks whether the server is running in qa environment.
func QA() bool {
	return Env() == "test" || Env() == "staging"
}

// Production checks whether the server is running in production environment.
func Production() bool {
	return Env() == "production"
}
