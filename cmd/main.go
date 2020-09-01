package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

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

	log.WithFields(log.Fields{
		"path": "unknown",
	}).Info("loading configuration")
	cobra.OnInitialize(initializeConfiguration)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	rootCmd.AddCommand(cmd.ServerCmd)

	rootCmd.Version = info.Version()
}

func initializeConfiguration() {
	config.LoadConfig(cfgFile)
}

func initializeLogger() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
