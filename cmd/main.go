package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/calvinverse/service.provisioning/internal/cmd"
	"github.com/calvinverse/service.provisioning/internal/config"
	"github.com/calvinverse/service.provisioning/internal/info"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use: "service.provisioning",
	}
)

func init() {
	initializeLogger()

	cobra.OnInitialize(initializeConfiguration)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	rootCmd.AddCommand(cmd.ServerCmd)

	rootCmd.Version = info.Version()
}

func initializeConfiguration() {
	log.Debug("Initializing. Loading configuration ...")
	config.LoadConfig(cfgFile)

	if viper.IsSet("log.level") {
		switch level := viper.GetString("log.level"); level {
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
