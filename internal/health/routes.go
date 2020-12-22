package health

import (
	"fmt"
	"net/http"
	"time"

	"github.com/calvinverse/service.provisioning.ui.web/internal/info"
	"github.com/calvinverse/service.provisioning.ui.web/internal/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// InfoResponse stores the response to an info request
type InfoResponse struct {
	BuildTime string `json:"buildtime"`
	Revision  string `json:"revision"`
	Version   string `json:"version"`
}

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

// Info godoc
// @Summary Respond to an info request
// @Description Respond to an info request with information about the application.
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} health.InfoResponse
// @Router /v1/self/info [get]
func (h *healthRouter) info(w http.ResponseWriter, r *http.Request) {
	response := InfoResponse{
		BuildTime: info.BuildTime(),
		Revision:  info.Revision(),
		Version:   info.Version(),
	}

	render.JSON(w, r, response)
}

// Ping godoc
// @Summary Respond to a ping request
// @Description Respond to a ping request with information about the application.
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} health.PingResponse
// @Router /v1/self/ping [get]
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
func (h *healthRouter) Routes(prefix string, r chi.Router) {
	r.Get(fmt.Sprintf("%s/ping", prefix), h.ping)
	r.Get(fmt.Sprintf("%s/info", prefix), h.info)
}

func (h *healthRouter) Version() int8 {
	return 1
}
