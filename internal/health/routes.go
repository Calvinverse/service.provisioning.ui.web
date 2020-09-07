package health

import (
	"fmt"
	"net/http"
	"time"

	"github.com/calvinverse/service.provisioning/internal/info"
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

// Routes creates the routes for the health package
func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", ping)

	return router
}

func ping(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	response := PingResponse{
		BuildTime: info.BuildTime(),
		Response:  fmt.Sprint("Pong - ", t.Format("Mon Jan _2 15:04:05 2006")),
		Revision:  info.Revision(),
		Version:   info.Version(),
	}

	render.JSON(w, r, response)
}
