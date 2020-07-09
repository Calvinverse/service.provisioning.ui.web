package cmd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/spf13/cobra"

	"github.com/calvinverse/service.provisioning/internal/health"
	"github.com/calvinverse/service.provisioning/internal/router"
	"github.com/calvinverse/service.provisioning/internal/web"
)

func executeServer(cmd *cobra.Command, args []string) {
	router := routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging error: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8080", router)) // NOTE: Get port from config
}

func routes() *chi.Mux {
	router := router.NewChiRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))

	// Based on this post and the comments: https://www.troyhunt.com/your-api-versioning-is-wrong-which-is/
	// Use the api/v1 approach
	//
	router.Route("api/v1", func(r chi.Router) {
		r.Mount("/self", health.Routes())
	})

	web.Routes(router)

	return Router
}

// ServerCmd is the server command
var ServerCmd = &cobra.Command{
	Use:   "Server",
	Short: "Runs the application as a server",
	Long:  "Runs the XXX application in server mode",
	Run:   executeServer,
}
