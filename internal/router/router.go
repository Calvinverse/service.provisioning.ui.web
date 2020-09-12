package router

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	"github.com/calvinverse/service.provisioning/internal/health"
	"github.com/calvinverse/service.provisioning/internal/web"
)

type RouterBuilder interface {
	NewChiRouter() *chi.Mux
}

type routerBuilder struct {
}

func (rb routerBuilder) apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}

func (rb routerBuilder) NewChiRouter() *chi.Mux {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}

	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Use(rb.newStructuredLogger(logger))

	router.Use(render.SetContentType(render.ContentTypeJSON))

	// Based on this post and the comments: https://www.troyhunt.com/your-api-versioning-is-wrong-which-is/
	// Use the api/v1 approach
	//
	router.Route("/api/v1", func(r chi.Router) {
		r.Use(rb.apiVersionCtx("v1"))
		r.Mount("/self", health.Routes())
	})

	web.Routes(router)

	return router
}

func (rb routerBuilder) newStructuredLogger(l *logrus.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructuredLogger{l})
}
