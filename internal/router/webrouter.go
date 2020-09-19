package router

import (
	"github.com/go-chi/chi"
)

// WebRouter provides a Chi Mux that contains a set of web routes.
type WebRouter interface {
	/// Routes provides the different routes for this WebRouter
	Routes(r *chi.Mux, rootRouter func() chi.Router)
}
