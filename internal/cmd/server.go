package cmd

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/calvinverse/service.provisioning/internal/router"
)

func executeServer(cmd *cobra.Command, args []string) {
	router := router.NewChiRouter()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	port := viper.GetInt("service.port")
	hostAddress := fmt.Sprintf(":%d", port)
	log.Debug(
		fmt.Sprintf(
			"Starting server on %s",
			hostAddress))
	log.Fatal(http.ListenAndServe(hostAddress, router))
}

// ServerCmd is the server command
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs the application as a server",
	Long:  "Runs the service.provisioning application in server mode",
	Run:   executeServer,
}
