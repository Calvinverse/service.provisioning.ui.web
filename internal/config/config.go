package config

import (
	"fmt"
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

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	// Only use consul if we have a host+port and consul key specified
	if viper.IsSet("consul.enabled") && viper.GetBool("consul.enabled") {
		loadFromConsul()
	}
}

func loadFromConsul() {

	viper.SetConfigType("yaml")

	consulHost := viper.GetString("consul.host")
	consulPort := viper.GetInt("consul.port")
	consulKeyPath := viper.GetString("consul.keyPath")
	if err := viper.AddRemoteProvider("consul", fmt.Sprintf("%s:%d", consulHost, consulPort), consulKeyPath); err != nil {
		log.Fatal(
			fmt.Sprintf(
				"Unable to connect to Consul at host %s:%d to read key %s. Error was %v",
				consulHost,
				consulPort,
				consulKeyPath,
				err))
	}

	if err := viper.ReadRemoteConfig(); err != nil {
		log.Warn(
			fmt.Sprintf(
				"Unable to read the configuration from Consul at key $s via host %s:%d. Error was %v",
				consulKeyPath,
				consulHost,
				consulPort,
				err))
	}

	go func() {
		for {
			time.Sleep(time.Second * 5) // delay after each request

			if err := viper.WatchRemoteConfig(); err != nil {
				log.Errorf("unable to read remote config: %v", err)
				continue
			}

			// unmarshal new config into our runtime config struct. you can also use channel
			// to implement a signal to notify the system of the changes
			viper.Unmarshal(&runtime_conf)
		}
	}()
}
