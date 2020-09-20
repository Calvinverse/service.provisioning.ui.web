package health

import (
	"fmt"
	"net/http"
	"time"

	"github.com/calvinverse/service.provisioning/internal/info"
	"github.com/calvinverse/service.provisioning/internal/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// PingResponse stores the response to a Ping request
type PingResponse struct {
	BuildTime string `json:"buildtime"`
	Response  string `json:"response"`
	Revision  string `json:"revision"`
	Version   string `json:"version"`
}

// NewHealthAPIRouter returns an APIRouter instance for the health routes.
func NewHealthAPIRouter() router.APIRouter {
	return &healthRouter{}
}

type healthRouter struct{}

// Ping godoc
// @Summary Respond to a ping request
// @Description Respond to a ping request with information about the application.
// @Accept  json
// @Produce  json
// @Success 200 {object} health.PingResponse
// @Router /api/v1/self/ping [get]
func (h *healthRouter) ping(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	response := PingResponse{
		BuildTime: info.BuildTime(),
		Response:  fmt.Sprint("Pong - ", t.Format("Mon Jan _2 15:04:05 2006")),
		Revision:  info.Revision(),
		Version:   info.Version(),
	}

	render.JSON(w, r, response)
}

func (h *healthRouter) Prefix() string {
	return "self"
}

// Routes creates the routes for the health package
func (h *healthRouter) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", h.ping)

	return router
}

func (h *healthRouter) Version() int8 {
	return 1
}
