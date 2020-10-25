package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

// Builder is used to create a new Chi Mux with all the routes and configurations set.
type Builder interface {
	New() *chi.Mux
}

// NewRouterBuilder creates a new instance of the Builder interface.
func NewRouterBuilder(
	apiRouters []APIRouter,
	webRouter WebRouter) Builder {
	return &routerBuilder{
		apiRouters: apiRouters,
		webRouter:  webRouter,
	}
}

type routerBuilder struct {
	apiRouters []APIRouter
	webRouter  WebRouter
}

func (rb routerBuilder) apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}

func (rb routerBuilder) New() *chi.Mux {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}

	router := rb.newChiRouter(logger)

	// Based on this post and the comments: https://www.troyhunt.com/your-api-versioning-is-wrong-which-is/
	// Use the api/v1 approach
	//
	router.Route("/api", func(r chi.Router) {
		for _, ar := range rb.apiRouters {
			r.Group(func(r chi.Router) {
				r.Use(rb.apiVersionCtx(
					fmt.Sprintf(
						"v%d",
						ar.Version())))

				prefix := fmt.Sprintf(
					"/v%d/%s",
					ar.Version(),
					ar.Prefix(),
				)
				ar.Routes(prefix, r)
			})
		}
	})

	rb.webRouter.Routes(router, func() chi.Router { return rb.newChiRouter(logger) })

	return router
}

func (rb routerBuilder) newChiRouter(logger *logrus.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Use(rb.newStructuredLogger(logger))

	router.Use(render.SetContentType(render.ContentTypeJSON))

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return router
}

func (rb routerBuilder) newStructuredLogger(l *logrus.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&structuredLogger{l})
}
