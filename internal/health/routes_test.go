package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoutes(t *testing.T) {
	request, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	router := Routes()
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
		t.Errorf("Handler returned wrong result: got %s wanted 'Pong'", actualResult.Response)
	}
}
