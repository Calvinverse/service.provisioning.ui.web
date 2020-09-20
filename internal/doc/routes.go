package doc

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/calvinverse/service.provisioning/internal/config"
	"github.com/calvinverse/service.provisioning/internal/router"
	"github.com/go-chi/chi"

	log "github.com/sirupsen/logrus"
)

// NewDocumentationRouter returns an APIRouter instance for the documentation routes.
func NewDocumentationRouter(config config.Configuration) router.APIRouter {
	return &docRouter{
		cfg: config,
	}
}

type docRouter struct {
	cfg config.Configuration
}

func (d *docRouter) Prefix() string {
	return "doc"
}

// Routes creates the routes for the doc package
func (d *docRouter) Routes() *chi.Mux {
	filesDir := ""
	if d.cfg.IsSet("doc.path") {
		filesDir = d.cfg.GetString("doc.path")
	} else {
		ex, err := os.Executable()
		if err != nil {
			log.Fatal(
				fmt.Sprintf(
					"Failed to locate the directory of the executable. Error was: %v",
					err))
		}

		workDir := filepath.Dir(ex)
		filesDir = filepath.Join(workDir, "doc")
	}

	log.Debug(
		fmt.Sprintf(
			"Using doc directory %s",
			filesDir))

	router := chi.NewRouter()

	// Load the index.html
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filesDir+"/openapi.json")
	})

	return router
}

func (d *docRouter) Version() int8 {
	return 1
}
