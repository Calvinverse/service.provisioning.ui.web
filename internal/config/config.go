package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/calvinverse/service.provisioning.ui.web/internal/observability"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"

	// Load viper/remote so that we can get configurations from Consul
	_ "github.com/spf13/viper/remote"
)

// Configuration defines the interface for configuration objects
type Configuration interface {
	GetInt(key string) int

	GetString(key string) string

	IsSet(key string) bool

	LoadConfiguration(cfgFile string) error
}

// NewConfiguration returns a new Configuration instance
func NewConfiguration() Configuration {
	return &concreteConfig{
		cfg: viper.New(),
	}
}

// concreteConfig implements the Configuration interface
type concreteConfig struct {
	cfg *viper.Viper
}

func (c *concreteConfig) GetInt(key string) int {
	return c.cfg.GetInt(key)
}

func (c *concreteConfig) GetString(key string) string {
	return c.cfg.GetString(key)
}

func (c *concreteConfig) IsSet(key string) bool {
	return c.cfg.IsSet(key)
}

// LoadConfiguration loads the configuration for the application from different configuration sources
func (c *concreteConfig) LoadConfiguration(cfgFile string) error {
	observability.LogDebug("Reading configuration ...")

	// From the environment
	c.cfg.SetEnvPrefix("PROVISION")
	c.cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	c.cfg.AutomaticEnv()

	if cfgFile != "" {
		observability.LogDebugWithFields(
			log.Fields{
				"configFile": cfgFile},
			"Reading configuration from file")

		c.cfg.SetConfigFile(cfgFile)
	}

	if err := c.cfg.ReadInConfig(); err != nil {
		observability.LogFatalWithFields(
			log.Fields{
				"error": err},
			"Configuration invalid")
		return err
	}

	// Only use consul if we have a host+port and consul key specified
	if c.cfg.IsSet("consul.enabled") && c.cfg.GetBool("consul.enabled") {
		if err := c.loadFromConsul(); err != nil {
			return err
		}
	}

	return nil
}

func (c *concreteConfig) loadFromConsul() error {

	c.cfg.SetConfigType("yaml")

	consulHost := c.GetString("consul.host")
	consulPort := c.GetInt("consul.port")
	consulKeyPath := c.GetString("consul.keyPath")
	observability.LogDebugWithFields(
		log.Fields{
			"consul_host":     consulHost,
			"consul_port":     consulPort,
			"consul_key_path": consulKeyPath},
		"Reading configuration from Consul")

	if err := c.cfg.AddRemoteProvider("consul", fmt.Sprintf("%s:%d", consulHost, consulPort), consulKeyPath); err != nil {
		observability.LogFatalWithFields(
			log.Fields{
				"consul_host":     consulHost,
				"consul_port":     consulPort,
				"consul_key_path": consulKeyPath,
				"error":           err},
			"Unable to connect to Consul")
		return err
	}

	if err := c.cfg.ReadRemoteConfig(); err != nil {
		observability.LogWarnWithFields(
			log.Fields{
				"consul_host":     consulHost,
				"consul_port":     consulPort,
				"consul_key_path": consulKeyPath,
				"error":           err},
			"Unable to read the configuration from Consul at the moment. Will continue to try.")

		// Don't return the error here because the inability to read from consul might be because the Consul
		// instance is currently not reachable. So we will continue to try.
	}

	go func() {
		for {
			time.Sleep(time.Second * 5) // delay after each request

			if err := c.cfg.WatchRemoteConfig(); err != nil {
				observability.LogWarnWithFields(
					log.Fields{
						"consul_host":     consulHost,
						"consul_port":     consulPort,
						"consul_key_path": consulKeyPath,
						"error":           err},
					"unable to read remote config")
				continue
			}

			observability.LogDebug("rereading remote config!")
			c.cfg.ReadRemoteConfig()
		}
	}()

	return nil
}
