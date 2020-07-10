package web

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/calvinverse/service.provisioning/internal/router"
)

var index []byte

func Routes(r *chi.Mux) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	workDir := filepath.Dir(ex)
	filesDir := filepath.Join(workDir, "client")

	// Load the index.html
	index, _ = ioutil.ReadFile(filesDir + "/index.html")

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filesDir+"/favicon.ico")
	})

	fileServer(r, "/client", http.Dir(filesDir))

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

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		w.Write(index)
	})

	return r
}
