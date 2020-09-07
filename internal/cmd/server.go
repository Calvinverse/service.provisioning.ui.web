package cmd

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/calvinverse/service.provisioning/internal/health"
	"github.com/calvinverse/service.provisioning/internal/router"
	"github.com/calvinverse/service.provisioning/internal/web"
)

func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}

func executeServer(cmd *cobra.Command, args []string) {
	router := routes()

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

func routes() *chi.Mux {
	router := router.NewChiRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))

	// Based on this post and the comments: https://www.troyhunt.com/your-api-versioning-is-wrong-which-is/
	// Use the api/v1 approach
	//
	router.Route("/api/v1", func(r chi.Router) {
		r.Use(apiVersionCtx("v1"))
		r.Mount("/self", health.Routes())
	})

	web.Routes(router)

	return router
}

// ServerCmd is the server command
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs the application as a server",
	Long:  "Runs the service.provisioning application in server mode",
	Run:   executeServer,
}
