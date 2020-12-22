package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/calvinverse/service.provisioning.ui.web/internal/info"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	request, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	instance := &healthRouter{}

	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		instance.Routes("", r)
	})

	router.ServeHTTP(w, request)

	actualResult := PingResponse{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	if !strings.HasPrefix(actualResult.Response, "Pong - ") {
		t.Errorf("Handler returned unexpected response: got %s wanted 'Pong'", actualResult.Response)
	}

	if actualResult.BuildTime != info.BuildTime() {
		t.Errorf("Handler returned unexpected build time: got %s wanted %s", actualResult.BuildTime, info.BuildTime())
	}

	if actualResult.Revision != info.Revision() {
		t.Errorf("Handler returned unexpected revision: got %s wanted %s", actualResult.Revision, info.Revision())
	}

	if actualResult.Version != info.Version() {
		t.Errorf("Handler returned unexpected build time: got %s wanted %s", actualResult.Version, info.Version())
	}
}
