package web

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/calvinverse/service.provisioning.ui.web/internal/config"
	"github.com/calvinverse/service.provisioning.ui.web/internal/observability"
	"github.com/calvinverse/service.provisioning.ui.web/internal/router"
	"github.com/go-chi/chi"
)

// NewWebRouter returns a new router.WebRouter instance
func NewWebRouter(config config.Configuration) router.WebRouter {
	return &webRouter{
		cfg: config,
	}
}

type webRouter struct {
	cfg config.Configuration
}

// Routes exports the web routes
func (w *webRouter) Routes(r *chi.Mux, rootRouter func() chi.Router) {

	filesDir := ""
	if w.cfg.IsSet("ui.path") {
		filesDir = w.cfg.GetString("ui.path")
		if _, err := os.Stat(filesDir); os.IsNotExist(err) {
			panic(
				fmt.Sprintf(
					"UI directory does not exist.Configuration was set to: %s",
					filesDir))
		}
	} else {
		ex, err := os.Executable()
		if err != nil {
			observability.LogFatalWithFields(
				log.Fields{
					"error": err},
				"Failed to locate the directory of the executable")
		}

		workDir := filepath.Dir(ex)
		filesDir = filepath.Join(workDir, "client")
	}

	observability.LogDebugWithFields(
		log.Fields{
			"directory": filesDir},
		"UI directory")

	// Load the index.html
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filesDir+"/index.html")
	})

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filesDir+"/favicon.ico")
	})

	w.fileServer(r, "/css", filesDir+"/css")
	w.fileServer(r, "/img", filesDir+"/img")
	w.fileServer(r, "/js", filesDir+"/js")
	w.fileServer(r, "/", filesDir)
}

func (w *webRouter) fileServer(r chi.Router, urlPath string, basePath string) {
	if strings.ContainsAny(urlPath, "{}*") {
		panic("FileServer does not permit any URL parameters")
	}

	root, _ := filepath.Abs(basePath)
	if _, err := os.Stat(root); os.IsNotExist(err) {
		panic("Static Documents Directory Not Found")
	}

	fs := http.StripPrefix(urlPath, http.FileServer(http.Dir(root)))

	if urlPath != "/" && urlPath[len(urlPath)-1] != '/' {
		r.Get(urlPath, http.RedirectHandler(urlPath+"/", 301).ServeHTTP)
		urlPath += "/"
	}

	r.Get(urlPath+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file := strings.Replace(r.RequestURI, urlPath, "/", 1)
		if _, err := os.Stat(root + file); os.IsNotExist(err) {
			http.ServeFile(w, r, path.Join(root, "index.html"))
			return
		}
		fs.ServeHTTP(w, r)
	}))
}
