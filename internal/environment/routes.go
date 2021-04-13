package environment

import (
	"net/http"

	"github.com/calvinverse/service.provisioning.ui.web/internal/router"
	"github.com/go-chi/chi"
)

func NewEnvironmentAPIRouter() router.APIRouter {
	return &environmentRouter{}
}

// Environment describes a single environment
type Environment struct{}

// https://www.google.com/search?client=firefox-b-d&q=golang+chi+get+query+params
// https://github.com/pressly/imgry/blob/master/server/server.go
// https://github.com/pressly/imgry/blob/master/server/middleware.go
// https://github.com/pressly/imgry/blob/master/server/handlers.go
// https://github.com/pressly/imgry/blob/bbb40ff8100ff84b8290005ebe080b7b07939372/server/middleware.go

type environmentRouter struct{}

// CreateEnvironment godoc
// @Summary Creates a new environment.
// @Description Creates a new environment based on the provided information.
// @Tags environment
// @Accept  json
// @Produce  json
// @Param id body environment.Environment true "Environment ID"
// @Success 201 {object} environment.Environment
// @Failure 404 {object} int
// @Failure 500 {object} int
// @Router /v1/environment [put]
func (h *environmentRouter) create(w http.ResponseWriter, r *http.Request) {
	//render.Status()
}

// DeleteEnvironment godoc
// @Summary Deletes an environment.
// @Description Deletes the environment with the given id.
// @Tags environment
// @Accept  json
// @Produce  json
// @Param id path string true "Environment ID"
// @Success 202 {object} environment.Environment
// @Failure 404 {object} int
// @Failure 500 {object} int
// @Router /v1/environment/{id} [delete]
func (h *environmentRouter) delete(w http.ResponseWriter, r *http.Request) {

}

// ShowEnvironment godoc
// @Summary Provide information about an environment.
// @Description Returns information about the environment with the given id.
// @Tags environment
// @Accept  json
// @Produce  json
// @Param id path string true "Environment ID"
// @Success 200 {object} environment.Environment
// @Failure 404 {object} int
// @Failure 500 {object} int
// @Router /v1/environment/{id} [get]
func (h *environmentRouter) get(w http.ResponseWriter, r *http.Request) {

}

// ListEnvironmentIDs godoc
// @Summary Provide the list of known environment IDs
// @Description Returns a list of known environment IDs.
// @Tags environment
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Failure 404 {object} int
// @Failure 500 {object} int
// @Router /v1/environment/ [get]
func (h *environmentRouter) list(w http.ResponseWriter, r *http.Request) {

}

func (h *environmentRouter) update(w http.ResponseWriter, r *http.Request) {

}

func (h *environmentRouter) Prefix() string {
	return "environment"
}

// Routes creates the routes for the health package
func (h *environmentRouter) Routes(prefix string, r chi.Router) {
	r.Route(prefix, func(r chi.Router) {
		r.Get("/", h.list)
		r.Get("/{id}", h.get)
		r.Delete("/{id}", h.delete)
		r.Put("/", h.create)
	})
}

func (h *environmentRouter) Version() int8 {
	return 1
}
