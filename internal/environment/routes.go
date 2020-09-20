package environment

import (
	"net/http"

	"github.com/calvinverse/service.provisioning/internal/router"
	"github.com/go-chi/chi"
)

func NewEnvironmentAPIRouter() router.APIRouter {
	return &environmentRouter{}
}

// https://www.google.com/search?client=firefox-b-d&q=golang+chi+get+query+params
// https://github.com/pressly/imgry/blob/master/server/server.go
// https://github.com/pressly/imgry/blob/master/server/middleware.go
// https://github.com/pressly/imgry/blob/master/server/handlers.go
// https://github.com/pressly/imgry/blob/bbb40ff8100ff84b8290005ebe080b7b07939372/server/middleware.go

type environmentRouter struct{}

func (h *environmentRouter) create(w http.ResponseWriter, r *http.Request) {

}

func (h *environmentRouter) delete(w http.ResponseWriter, r *http.Request) {

}

func (h *environmentRouter) get(w http.ResponseWriter, r *http.Request) {

}

func (h *environmentRouter) list(w http.ResponseWriter, r *http.Request) {

}

func (h *environmentRouter) update(w http.ResponseWriter, r *http.Request) {

}

func (h *environmentRouter) Prefix() string {
	return "environment"
}

// Routes creates the routes for the health package
func (h *environmentRouter) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/", h.list)
		r.Get("/{id}", h.get)
		r.Delete("/{id}", h.delete)
		r.Put("/", h.create)
	})

	return router
}

func (h *environmentRouter) Version() int8 {
	return 1
}
