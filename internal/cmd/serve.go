package cmd

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/spf13/cobra"

	"github.com/calvinverse/service.provisioning.ui.web/internal/config"
	info "github.com/calvinverse/service.provisioning.ui.web/internal/info"
	"github.com/calvinverse/service.provisioning.ui.web/internal/observability"
	"github.com/calvinverse/service.provisioning.ui.web/internal/router"
)

// ServeCommandBuilder creates new Cobra Commands for the server capability.
type ServeCommandBuilder interface {
	New() *cobra.Command
}

// @title Service.Provisioning.UI.Web server API
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
// NewServeCommandBuilder creates a new instance of the ServerCommandBuilder interface.
func NewServeCommandBuilder(config config.Configuration, builder router.Builder) ServeCommandBuilder {
	return &serveCommandBuilder{
		cfg:     config,
		builder: builder,
	}
}

type serveCommandBuilder struct {
	cfg     config.Configuration
	builder router.Builder
}

func (s serveCommandBuilder) New() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Runs the application as a server",
		Long:  "Runs the service.provisioning.ui.web application in server mode",
		RunE:  s.executeServer,
	}
}

func (s *serveCommandBuilder) configureHealthCheck() error {
	check := ServeLivelinessCheck()

	center := info.GetHealthCenter()
	err := center.RegisterLivelinessCheck(
		check,
		30*time.Second,
		5*time.Second,
		false,
	)

	return err
}

func (s *serveCommandBuilder) createRouter() (*chi.Mux, error) {
	router := s.builder.New()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		observability.LogInfoWithFields(log.Fields{
			"method": method,
			"route":  route},
			"Request received")
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		observability.LogPanicWithFields(log.Fields{
			"error": err.Error()},
			"An error occurred")
		return nil, err
	}
	return router, nil
}

func (s serveCommandBuilder) executeServer(cmd *cobra.Command, args []string) error {
	router, err := s.createRouter()
	if err != nil {
		observability.LogFatal(err)
		return err
	}

	err = s.configureHealthCheck()
	if err != nil {
		observability.LogFatal(err)
		return err
	}

	host, port := s.getHostConnectionDetails()
	observability.LogDebugWithFields(log.Fields{
		"host": host,
		"port": port},
		"Starting server")
	hostAddress := fmt.Sprintf("%s:%d", host, port)
	if err := http.ListenAndServe(hostAddress, router); err != nil {
		observability.LogFatal(err)
		return err
	}

	return nil
}

func (s *serveCommandBuilder) getHostConnectionDetails() (string, int) {
	port := s.cfg.GetInt("service.port")
	return "", port
}
