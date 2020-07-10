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
	cobra.OnInitialize(config.LoadConfig)

	rootCmd.AddCommand(cmd.ServerCmd)

	rootCmd.Version = info.Version()
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
