package web

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/calvinverse/service.provisioning/internal/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/spf13/viper"
)

// Routes exports the web routes
func Routes(r *chi.Mux) {
	workDir := viper.GetString("config.ui.path")
	filesDir := filepath.Join(workDir, "client")

	// Load the index.html
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filesDir+"/index.html")
	})

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filesDir+"/favicon.ico")
	})

	fileServer(r, "/css", http.Dir(filesDir+"/css"))
	fileServer(r, "/img", http.Dir(filesDir+"/img"))
	fileServer(r, "/js", http.Dir(filesDir+"/js"))

	r.Mount("/", rootRouter())
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		destination := path + "/"
		r.Get(path, http.RedirectHandler(destination, 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func rootRouter() chi.Router {
	r := router.NewChiRouter()

	r.Use(middleware.NoCache)

	return r
}
