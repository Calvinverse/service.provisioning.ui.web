package router

import (
	"github.com/go-chi/chi"
)

// APIRouter provides a Chi Mux that contains a set of API routes.
type APIRouter interface {
	// Prefix is the prefix of the API method, e.g. self or template.
	Prefix() string

	/// Routes provides the different routes for this APIRouter
	Routes(prefix string, r chi.Router)

	// Version returns the API version for the routes.
	Version() int8
}
