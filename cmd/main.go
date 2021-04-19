package main

import (
	"github.com/spf13/cobra"

	"github.com/calvinverse/service.provisioning.ui.web/internal/config"
	"github.com/calvinverse/service.provisioning.ui.web/internal/meta"
	"github.com/calvinverse/service.provisioning.ui.web/internal/observability"
	"github.com/calvinverse/service.provisioning.ui.web/internal/service"
)

var (
	cfgFile string

	cfg config.Configuration

	resolver service.Resolver

	rootCmd = &cobra.Command{
		Use:     "server",
		Version: meta.Version(),
	}
)

func init() {
	observability.InitializeLogger()

	cfg = config.NewConfiguration()
	resolver := service.NewResolver(cfg)

	cobra.OnInitialize(initializeConfiguration)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	for _, subCommand := range resolver.ResolveCommands() {
		rootCmd.AddCommand(subCommand)
	}
}

func initializeConfiguration() {
	observability.LogDebug("Initializing. Loading configuration ...")
	cfg.LoadConfiguration(cfgFile)

	if cfg.IsSet("log.level") {
		level := cfg.GetString("log.level")
		observability.SetLogLevel(level)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		observability.LogFatal(err)
	}
}
