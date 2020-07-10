package config

import (
	"strings"

	"github.com/spf13/viper"
)

// LoadConfig loads the configuration for the application from different configuration sources
func LoadConfig() {

	// From the environment
	viper.SetEnvPrefix("PROVISION")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// From a config file
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	// From Consul

	if err := viper.ReadInConfig(); err != nil {
		// Write something to the log
	}
}
