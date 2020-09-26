package cmd

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/spf13/cobra"

	"github.com/calvinverse/service.provisioning/internal/config"
	"github.com/calvinverse/service.provisioning/internal/router"
)

// ServerCommandBuilder creates new Cobra Commands for the server capability.
type ServerCommandBuilder interface {
	New() *cobra.Command
}

// @title Service.Provisioning server API
// @version 1.0
// @description Provides information about deployed environments and the templates used to created these environments.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

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
//
// NewCommandBuilder creates a new instance of the ServerCommandBuilder interface.
func NewCommandBuilder(config config.Configuration, builder router.Builder) ServerCommandBuilder {
	return &serverCommandBuilder{
		cfg:     config,
		builder: builder,
	}
}

type serverCommandBuilder struct {
	cfg     config.Configuration
	builder router.Builder
}

func (s serverCommandBuilder) New() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Runs the application as a server",
		Long:  "Runs the service.provisioning application in server mode",
		RunE:  s.executeServer,
	}
}

func (s serverCommandBuilder) executeServer(cmd *cobra.Command, args []string) error {
	router := s.builder.New()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
		return err
	}

	port := s.cfg.GetInt("service.port")
	hostAddress := fmt.Sprintf(":%d", port)
	log.Debug(
		fmt.Sprintf(
			"Starting server on %s",
			hostAddress))
	if err := http.ListenAndServe(hostAddress, router); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
