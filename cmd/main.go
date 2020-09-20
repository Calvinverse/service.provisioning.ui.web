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

	cfg = config.NewConfiguration()
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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
