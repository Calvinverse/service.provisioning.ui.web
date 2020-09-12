package router

import (
	"github.com/go-chi/chi"
)

type ApiRouter interface {
	ApiRoutes() *chi.Mux

	ApiVersion() int8
}
