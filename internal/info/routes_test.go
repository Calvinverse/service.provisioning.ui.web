package info

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi"
)

//
// Mocks
//

type mockHealthService struct {
	liveliness *HealthStatus
	readiness  *HealthStatus
	started    *StartedStatus
	error      error
}

func (h *mockHealthService) Liveliness() (*HealthStatus, error) {
	return h.liveliness, h.error
}

func (h *mockHealthService) Readiness() (*HealthStatus, error) {
	return h.readiness, h.error
}

func (h *mockHealthService) Started() (*StartedStatus, error) {
	return h.started, h.error
}

type mockError struct{}

func (e *mockError) Error() string {
	return "some text"
}

//
// Info
//

func TestInfoWithAcceptHeaderSetToJson(t *testing.T) {
	request := setupRequest("/info", "application/json", make(map[string]string))

	w := httptest.NewRecorder()

	router := setupHTTPRouter(&selfRouter{})
	router.ServeHTTP(w, request)

	validateInfoWithAcceptHeader(t, w, decodeJSONFromResponseBody)
}

func TestInfoWithAcceptHeaderSetToXml(t *testing.T) {
	request := setupRequest("/info", "application/xml", make(map[string]string))

	w := httptest.NewRecorder()

	router := setupHTTPRouter(&selfRouter{})
	router.ServeHTTP(w, request)

	validateInfoWithAcceptHeader(t, w, decodeXMLFromResponseBody)
}

func TestInfoWithoutHeader(t *testing.T) {
	request := setupRequest("/info", "", make(map[string]string))

	w := httptest.NewRecorder()

	router := setupHTTPRouter(&selfRouter{})
	router.ServeHTTP(w, request)

	validateWithoutAcceptHeader(t, w, decodeJSONFromResponseBody)
}

//
// liveliness
//

func TestLivelinessWithFailingHealthAndHeaderSetToJson(t *testing.T) {
	request, _ := http.NewRequest("GET", "/liveliness", nil)
	request.Header.Set("Accept", "application/json")

	w := httptest.NewRecorder()

	healthService := &mockHealthService{
		error: &mockError{},
	}
	instance := &selfRouter{
		healthService: healthService,
	}

	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		instance.Routes("", r)
	})

	router.ServeHTTP(w, request)

	actualResult := livelinessSummaryResponse{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	if status := w.Code; status != http.StatusInternalServerError {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusInternalServerError)
	}

	if actualResult.Status != Failed {
		t.Errorf("Handler returned unexpected status: got %s wanted %s", actualResult.Status, Failed)
	}

	if len(actualResult.Checks) != 0 {
		t.Errorf("Handler returned unexpected number of checks: got %d wanted %d", len(actualResult.Checks), 0)
	}
}

func TestLivelinessWithNoAccept(t *testing.T) {
	request := setupRequest("/liveliness", "", make(map[string]string))

	w := httptest.NewRecorder()

	healthService := &mockHealthService{
		error: &mockError{},
	}
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	validateWithoutAcceptHeader(t, w, decodeJSONFromResponseBody)
}

func TestLivelinessDetailedWithAcceptHeaderSetToJson(t *testing.T) {
	queryParameters := map[string]string{
		"type": "detailed",
	}

	request := setupRequest("/liveliness", "application/json", queryParameters)

	w := httptest.NewRecorder()

	numberOfChecks := 2

	healthService := createHealthServiceWithChecks(numberOfChecks, 0)
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	actualResult := livelinessDetailedResponse{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	validateLivelinessDetailedResponse(t, numberOfChecks, actualResult)
}

func TestLivelinessDetailedWithAcceptHeaderSetToXml(t *testing.T) {
	queryParameters := map[string]string{
		"type": "detailed",
	}

	request := setupRequest("/liveliness", "application/xml", queryParameters)

	w := httptest.NewRecorder()

	numberOfChecks := 2

	healthService := createHealthServiceWithChecks(numberOfChecks, 0)
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	actualResult := livelinessDetailedResponse{}
	xml.NewDecoder(w.Body).Decode(&actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	validateLivelinessDetailedResponse(t, numberOfChecks, actualResult)
}

func TestLivelinessSummaryWithAcceptHeaderSetToJson(t *testing.T) {
	queryParameters := map[string]string{
		"type": "summary",
	}

	request := setupRequest("/liveliness", "application/json", queryParameters)

	w := httptest.NewRecorder()

	numberOfChecks := 2

	healthService := createHealthServiceWithChecks(numberOfChecks, 0)
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	actualResult := livelinessSummaryResponse{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	validateLivelinessSummaryResponse(t, numberOfChecks, actualResult)
}

func TestLivelinessSummaryWithAcceptHeaderSetToXml(t *testing.T) {
	queryParameters := map[string]string{
		"type": "summary",
	}

	request := setupRequest("/liveliness", "application/xml", queryParameters)

	w := httptest.NewRecorder()

	numberOfChecks := 2

	healthService := createHealthServiceWithChecks(numberOfChecks, 0)
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	actualResult := livelinessSummaryResponse{}
	xml.NewDecoder(w.Body).Decode(&actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	validateLivelinessSummaryResponse(t, numberOfChecks, actualResult)
}

//
// ping
//

func TestPingyWithAcceptHeaderSetToJson(t *testing.T) {
	request := setupRequest("/ping", "application/json", make(map[string]string))

	w := httptest.NewRecorder()

	router := setupHTTPRouter(&selfRouter{})
	router.ServeHTTP(w, request)

	validatePingWithAcceptHeader(t, w, decodeJSONFromResponseBody)
}

func TestPingyWithAcceptHeaderSetToXml(t *testing.T) {
	request := setupRequest("/ping", "application/xml", make(map[string]string))

	w := httptest.NewRecorder()

	router := setupHTTPRouter(&selfRouter{})
	router.ServeHTTP(w, request)

	validatePingWithAcceptHeader(t, w, decodeXMLFromResponseBody)
}

func TestPingWithoutHeader(t *testing.T) {
	request := setupRequest("/ping", "", make(map[string]string))

	w := httptest.NewRecorder()

	router := setupHTTPRouter(&selfRouter{})
	router.ServeHTTP(w, request)

	validateWithoutAcceptHeader(t, w, decodeJSONFromResponseBody)
}

//
// readiness
//

func TestReadinessWithAcceptHeaderSetToJson(t *testing.T) {
	request := setupRequest("/readiness", "application/json", make(map[string]string))

	w := httptest.NewRecorder()

	numberOfChecks := 2

	healthService := createHealthServiceWithChecks(0, numberOfChecks)
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	actualResult := readinessResponse{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	validateReadinessResponse(t, numberOfChecks, actualResult)
}

func TestReadinessWithAcceptHeaderSetToXml(t *testing.T) {
	request := setupRequest("/readiness", "application/xml", make(map[string]string))

	w := httptest.NewRecorder()

	numberOfChecks := 2

	healthService := createHealthServiceWithChecks(0, numberOfChecks)
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	actualResult := readinessResponse{}
	xml.NewDecoder(w.Body).Decode(&actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	validateReadinessResponse(t, numberOfChecks, actualResult)
}

func TestReadinessWithNoAccept(t *testing.T) {
	request := setupRequest("/readiness", "", make(map[string]string))

	w := httptest.NewRecorder()

	healthService := &mockHealthService{
		error: &mockError{},
	}
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	validateWithoutAcceptHeader(t, w, decodeJSONFromResponseBody)
}

//
// started
//

func TestStartedWithAcceptHeaderSetToJson(t *testing.T) {
	request := setupRequest("/started", "application/json", make(map[string]string))

	w := httptest.NewRecorder()

	numberOfChecks := 2

	healthService := createHealthServiceWithChecks(0, numberOfChecks)
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	actualResult := startedResponse{}
	json.NewDecoder(w.Body).Decode(&actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	validateStartedResponse(t, actualResult)
}

func TestStartedWithAcceptHeaderSetToXml(t *testing.T) {
	request := setupRequest("/started", "application/xml", make(map[string]string))

	w := httptest.NewRecorder()

	numberOfChecks := 2

	healthService := createHealthServiceWithChecks(0, numberOfChecks)
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	actualResult := startedResponse{}
	xml.NewDecoder(w.Body).Decode(&actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	validateStartedResponse(t, actualResult)
}

func TestStartedWithNoAccept(t *testing.T) {
	request := setupRequest("/started", "", make(map[string]string))

	w := httptest.NewRecorder()

	healthService := &mockHealthService{
		error: &mockError{},
	}
	instance := &selfRouter{
		healthService: healthService,
	}

	router := setupHTTPRouter(instance)
	router.ServeHTTP(w, request)

	validateWithoutAcceptHeader(t, w, decodeJSONFromResponseBody)
}

//
// Helper functions
//

type decodeResponseBody func(buffer *bytes.Buffer, v interface{}) error

func decodeJSONFromResponseBody(buffer *bytes.Buffer, v interface{}) error {
	return json.NewDecoder(buffer).Decode(v)
}

func decodeXMLFromResponseBody(buffer *bytes.Buffer, v interface{}) error {
	return xml.NewDecoder(buffer).Decode(v)
}

type validateResponse func(t *testing.T, w *httptest.ResponseRecorder, decode decodeResponseBody)

//
// Setup functions
//

func createHealthServiceWithChecks(numberOfLivelinessChecks int, numberOfReadinessChecks int) *mockHealthService {
	livelinessStatus := createLivelinessStatus(numberOfLivelinessChecks)
	readinessStatus := createReadinessStatus(numberOfReadinessChecks)
	startedStatus := &StartedStatus{
		Timestamp: time.Date(2021, time.January, 1, 1, 1, 1, 0, time.Local),
	}

	healthService := &mockHealthService{
		liveliness: livelinessStatus,
		readiness:  readinessStatus,
		started:    startedStatus,
		error:      nil,
	}
	return healthService
}

func createLivelinessStatus(numberOfChecks int) *HealthStatus {
	var checks []HealthCheckResult
	for i := 0; i < numberOfChecks; i++ {
		check := HealthCheckResult{
			IsSuccess: true,
			Name:      strconv.Itoa(i),
			Timestamp: time.Date(2021, time.January, i, i, i, i, 0, time.Local),
		}
		checks = append(checks, check)
	}

	status := &HealthStatus{
		Checks:    checks,
		IsHealthy: true,
	}
	return status
}

func createReadinessStatus(numberOfChecks int) *HealthStatus {
	var checks []HealthCheckResult
	for i := 0; i < numberOfChecks; i++ {
		check := HealthCheckResult{
			IsSuccess: true,
			Name:      strconv.Itoa(i),
			Timestamp: time.Date(2021, time.January, i, i, i, i, 0, time.Local),
		}
		checks = append(checks, check)
	}

	status := &HealthStatus{
		Checks:    checks,
		IsHealthy: true,
	}
	return status
}

func setupHTTPRouter(instance *selfRouter) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/", func(r chi.Router) {
		instance.Routes("", r)
	})

	return router
}

func setupRequest(path string, acceptHeader string, queryItems map[string]string) *http.Request {
	request, _ := http.NewRequest("GET", path, nil)

	if acceptHeader != "" {
		request.Header.Set("Accept", acceptHeader)
	}

	for k, v := range queryItems {
		q := request.URL.Query()
		q.Add(k, v)
		request.URL.RawQuery = q.Encode()
	}

	return request
}

//
// Validation functions
//

func validateWithoutAcceptHeader(t *testing.T, w *httptest.ResponseRecorder, decode decodeResponseBody) {
	if status := w.Code; status != http.StatusUnsupportedMediaType {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusUnsupportedMediaType)
	}
}

func validateInfoWithAcceptHeader(t *testing.T, w *httptest.ResponseRecorder, decode decodeResponseBody) {
	actualResult := infoResponse{}
	decode(w.Body, &actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	if actualResult.BuildTime != BuildTime() {
		t.Errorf("Handler returned unexpected build time: got %s wanted %s", actualResult.BuildTime, BuildTime())
	}

	if actualResult.Revision != Revision() {
		t.Errorf("Handler returned unexpected revision: got %s wanted %s", actualResult.Revision, Revision())
	}

	if actualResult.Version != Version() {
		t.Errorf("Handler returned unexpected build time: got %s wanted %s", actualResult.Version, Version())
	}
}

func validatePingWithAcceptHeader(t *testing.T, w *httptest.ResponseRecorder, decode decodeResponseBody) {
	actualResult := pingResponse{}
	decode(w.Body, &actualResult)

	if status := w.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK)
	}

	if !strings.HasPrefix(actualResult.Response, "Pong - ") {
		t.Errorf("Handler returned unexpected response: got %s wanted 'Pong'", actualResult.Response)
	}
}

func validateLivelinessDetailedResponse(t *testing.T, expectedNumberOfChecks int, result livelinessDetailedResponse) {
	if result.Status != Success {
		t.Errorf("Handler returned unexpected status: got %s wanted %s", result.Status, Success)
	}

	if len(result.Checks) != expectedNumberOfChecks {
		t.Errorf("Handler returned unexpected number of checks: got %d wanted %d", len(result.Checks), expectedNumberOfChecks)
	}

	for i, c := range result.Checks {
		if c.Name != strconv.Itoa(i) {
			t.Errorf("Check %d has an unexpected name: got %s wanted %s", i, c.Name, strconv.Itoa(i))
		}

		if c.Status != Success {
			t.Errorf("Check %d had an unexpected status. Expected Success got %s", i, c.Status)
		}

		parsedTime, err := time.Parse(time.RFC3339, c.Timestamp)
		if err != nil {
			t.Errorf("Check %d contained a timestamp that was not parsable. Got %s", i, c.Timestamp)
		}

		expectedTime := time.Date(2021, time.January, i, i, i, i, 0, time.Local)
		if !parsedTime.Equal(expectedTime) {
			t.Errorf("Check %d had an unexpected timestamp. Got %s wanted %s", i, c.Timestamp, expectedTime.Format(time.RFC3339))
		}
	}
}

func validateLivelinessSummaryResponse(t *testing.T, expectedNumberOfChecks int, result livelinessSummaryResponse) {
	if result.Status != Success {
		t.Errorf("Handler returned unexpected status: got %s wanted %s", result.Status, Success)
	}

	if len(result.Checks) != expectedNumberOfChecks {
		t.Errorf("Handler returned unexpected number of checks: got %d wanted %d", len(result.Checks), expectedNumberOfChecks)
	}

	for i, k := range result.Checks {

		expectedName := strconv.Itoa(i)
		if k.Name != expectedName {
			t.Errorf("Check has an unexpected name: got %s wanted %s", k.Name, expectedName)
		}

		if k.Status != Success {
			t.Errorf("Check had an unexpected status. Expected Success got %s", k.Status)
		}
	}
}

func validateReadinessResponse(t *testing.T, expectedNumberOfChecks int, result readinessResponse) {
	if result.Status != Success {
		t.Errorf("Handler returned unexpected status: got %s wanted %s", result.Status, Success)
	}

	if len(result.Checks) != expectedNumberOfChecks {
		t.Errorf("Handler returned unexpected number of checks: got %d wanted %d", len(result.Checks), expectedNumberOfChecks)
	}

	for i, k := range result.Checks {

		expectedName := strconv.Itoa(i)
		if k.Name != expectedName {
			t.Errorf("Check has an unexpected name: got %s wanted %s", k.Name, expectedName)
		}

		if k.Status != Success {
			t.Errorf("Check had an unexpected status. Expected Success got %s", k.Status)
		}
	}
}

func validateStartedResponse(t *testing.T, result startedResponse) {
	parsedTime, err := time.Parse(time.RFC3339, result.Timestamp)
	if err != nil {
		t.Errorf("Check contained a timestamp that was not parsable. Got %s", result.Timestamp)
	}

	expectedTime := time.Date(2021, time.January, 1, 1, 1, 1, 0, time.Local)
	if !parsedTime.Equal(expectedTime) {
		t.Errorf("Check had an unexpected timestamp. Got %s wanted %s", parsedTime.String(), expectedTime.String())
	}
}
