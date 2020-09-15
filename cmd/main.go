package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/calvinverse/service.provisioning/internal/config"
	"github.com/calvinverse/service.provisioning/internal/info"
	"github.com/calvinverse/service.provisioning/internal/service"
)

var (
	cfgFile string

	cfg config.Configuration

	resolver service.Resolver

	rootCmd = &cobra.Command{
		Use:     "service.provisioning",
		Version: info.Version(),
	}
)

func init() {
	initializeLogger()

	cfg := config.NewConfiguration()
	resolver := service.NewResolver(cfg)

	cobra.OnInitialize(initializeConfiguration)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	for _, subCommand := range resolver.ResolveCommands() {
		rootCmd.AddCommand(subCommand)
	}
}

func initializeConfiguration() {
	log.Debug("Initializing. Loading configuration ...")
	cfg.LoadConfiguration(cfgFile)

	if cfg.IsSet("log.level") {
		switch level := cfg.GetString("log.level"); level {
		case "trace":
			log.SetLevel(log.TraceLevel)
		case "debug":
			log.SetLevel(log.DebugLevel)
		case "info":
			log.SetLevel(log.InfoLevel)
		case "warn":
			log.SetLevel(log.WarnLevel)
		case "error":
			log.SetLevel(log.ErrorLevel)
		case "fatal":
			log.SetLevel(log.FatalLevel)
		default:
			log.SetLevel(log.InfoLevel)
		}
	}
}

func initializeLogger() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.InfoLevel)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
