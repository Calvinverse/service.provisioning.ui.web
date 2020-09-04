package config

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"

	// Load viper/remote so that we can get configurations from Consul
	_ "github.com/spf13/viper/remote"
)

// LoadConfig loads the configuration for the application from different configuration sources
func LoadConfig(cfgFile string) {

	// From the environment
	viper.SetEnvPrefix("PROVISION")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	// Only use consul if we have a host+port and consul key specified

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func loadFromConsul() {
	// From Consul
	viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
	viper.SetConfigType("yaml")

	go func() {
		for {
			time.Sleep(time.Second * 5) // delay after each request

			// currently, only tested with etcd support
			err := viper.WatchRemoteConfig()
			if err != nil {
				log.Errorf("unable to read remote config: %v", err)
				continue
			}

			// unmarshal new config into our runtime config struct. you can also use channel
			// to implement a signal to notify the system of the changes
			viper.Unmarshal(&runtime_conf)
		}
	}()
}
