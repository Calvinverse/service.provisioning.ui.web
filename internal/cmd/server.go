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
		Run:   s.executeServer,
	}
}

func (s serverCommandBuilder) executeServer(cmd *cobra.Command, args []string) {
	router := s.builder.New()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	port := s.cfg.GetInt("service.port")
	hostAddress := fmt.Sprintf(":%d", port)
	log.Debug(
		fmt.Sprintf(
			"Starting server on %s",
			hostAddress))
	log.Fatal(http.ListenAndServe(hostAddress, router))
}
