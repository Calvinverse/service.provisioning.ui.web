package info

import (
	"fmt"
	"mime"
	"net/http"
	"time"

	"github.com/calvinverse/service.provisioning.ui.web/internal/meta"
	"github.com/calvinverse/service.provisioning.ui.web/internal/observability"
	"github.com/calvinverse/service.provisioning.ui.web/internal/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	log "github.com/sirupsen/logrus"
)

// DetailedCheckInformation stores information about the status of a health check.
type detailedCheckInformation struct {
	// Description returns the description of the health check status.
	Description string `json:"description"`

	// Name returns the name of the health check.
	Name string `json:"name"`

	// Status returns the status of the health check, either success or failure.
	Status string `json:"status"`

	// Timestamp returns the time the healtcheck was executed.
	Timestamp string `json:"timestamp"`
}

// InfoResponse stores the response to an info request
type infoResponse struct {
	// BuildTime stores the date and time the application was built.
	BuildTime string `json:"buildtime"`

	// Revision stores the GIT SHA of the commit on which the application build was based.
	Revision string `json:"revision"`

	// Version stores the version number of the application.
	Version string `json:"version"`
}

// LivelinessDetailedResponse stores detailed information about the liveliness of the application, indicating if the application is healthy
type livelinessDetailedResponse struct {
	// Status of all the health checks
	Checks []detailedCheckInformation `json:"checks"`

	// Global status
	Status string `json:"status"`

	// Timestamp the liveliness response was created at
	Timestamp string `json:"time"`
}

// LivelinessSummaryResponse stores condensed information about the liveliness of the application, indicating if the application is healthy
type livelinessSummaryResponse struct {
	// Status of all health checks
	Checks []summaryCheckInformation `json:"checks"`

	// Global status
	Status string `json:"status"`

	// Timestamp the liveliness response was created at
	Timestamp string `json:"time"`
}

// PingResponse stores the response to a Ping request
type pingResponse struct {
	Response string `json:"response"`
}

// ReadinessResponse stores information about the readiness of the application, indicating whether the application is ready to serve responses.
type readinessResponse struct {
	// Status of all health checks
	Checks []summaryCheckInformation `json:"checks"`

	// Status returns the status of the readiness check, either success or failure.
	Status string `json:"status"`

	// Timestamp returns the timestamp at which the last readiness check was executed.
	Timestamp string `json:"time"`
}

// StartedResponse stores information about the starting state of the application, indicating whether the application has started successfully.
type startedResponse struct {
	Timestamp string `json:"time"`
}

// SummaryCheckInformation stores the minimal information about the status of a health check.
type summaryCheckInformation struct {
	// Name returns the name of the health check.
	Name string `json:"name"`

	// Status returns the status of the health check, either success or failure.
	Status string `json:"status"`
}

// NewSelfAPIRouter returns an APIRouter instance for the health routes.
func NewSelfAPIRouter() router.APIRouter {
	return &selfRouter{
		healthService: GetStatusReporter(),
	}
}

// selfRouter defines an APIRouter that routes the 'self' metadata routes.
type selfRouter struct {
	healthService StatusReporter
}

func (h *selfRouter) Prefix() string {
	return "self"
}

// Routes creates the routes for the health package
func (h *selfRouter) Routes(prefix string, r chi.Router) {
	r.Get(fmt.Sprintf("%s/info", prefix), h.info)
	r.Get(fmt.Sprintf("%s/liveliness", prefix), h.liveliness)
	r.Get(fmt.Sprintf("%s/ping", prefix), h.ping)
	r.Get(fmt.Sprintf("%s/readiness", prefix), h.readiness)
	r.Get(fmt.Sprintf("%s/started", prefix), h.started)
}

func (h *selfRouter) Version() int8 {
	return 1
}

// Info godoc
// @Summary Respond to an info request
// @Description Respond to an info request with information about the application.
// @Tags health
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Success 200 {object} info.infoResponse
// @Failure 415 {string} string "Unsupported media type"
// @Router /v1/self/info [get]
func (h *selfRouter) info(w http.ResponseWriter, r *http.Request) {
	response := infoResponse{
		BuildTime: meta.BuildTime(),
		Revision:  meta.Revision(),
		Version:   meta.Version(),
	}

	h.responseBody(w, r, http.StatusOK, response)
}

// Liveliness godoc
// @Summary Respond to an liveliness request
// @Description Respond to an liveliness request with information about the status of the latest health checks.
// @Tags health
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Param type query string false "options are summary or detailed" Enums(summary, detailed)
// @Success 200 {object} info.livelinessDetailedResponse
// @Failure 415 {string} string "Unsupported media type"
// @Router /v1/self/liveliness [get]
func (h *selfRouter) liveliness(w http.ResponseWriter, r *http.Request) {
	healthStatus, err := h.healthService.Liveliness()
	if err != nil {
		healthStatus = &HealthStatus{
			Checks:    make([]HealthCheckResult, 0, 0),
			IsHealthy: false,
		}

		h.livelinessSummaryResponse(w, r, healthStatus)
		return
	}

	responseType := r.URL.Query().Get("type")
	switch responseType {
	case "detailed":
		h.livelinessDetailedResponse(w, r, healthStatus)
	case "summary":
		fallthrough
	default:
		h.livelinessSummaryResponse(w, r, healthStatus)
	}
}

func (h *selfRouter) livelinessDetailedResponse(w http.ResponseWriter, r *http.Request, status *HealthStatus) {
	t := time.Now()

	statusText := statusToText(status.IsHealthy)
	responseCode := statusToResponseCode(status.IsHealthy)

	var checkResults []detailedCheckInformation
	checkResults = make([]detailedCheckInformation, 0, len(status.Checks))
	for _, check := range status.Checks {

		result := detailedCheckInformation{
			Description: check.Description,
			Name:        check.Name,
			Status:      statusToText(check.IsSuccess),
			Timestamp:   check.Timestamp.Format(time.RFC3339),
		}
		checkResults = append(checkResults, result)
	}

	response := &livelinessDetailedResponse{
		Checks:    checkResults,
		Status:    statusText,
		Timestamp: t.Format(time.RFC3339),
	}

	h.responseBody(w, r, responseCode, response)
}

func (h *selfRouter) livelinessSummaryResponse(w http.ResponseWriter, r *http.Request, status *HealthStatus) {
	t := time.Now()

	statusText := statusToText(status.IsHealthy)
	responseCode := statusToResponseCode(status.IsHealthy)

	var checkResults []summaryCheckInformation
	checkResults = make([]summaryCheckInformation, 0, len(status.Checks))
	for _, check := range status.Checks {

		result := summaryCheckInformation{
			Name:   check.Name,
			Status: statusToText(check.IsSuccess),
		}
		checkResults = append(checkResults, result)
	}

	response := &livelinessSummaryResponse{
		Checks:    checkResults,
		Status:    statusText,
		Timestamp: t.Format(time.RFC3339),
	}

	h.responseBody(w, r, responseCode, response)
}

// Ping godoc
// @Summary Respond to a ping request
// @Description Respond to a ping request with a pong response.
// @Tags health
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Success 200 {object} info.pingResponse
// @Failure 415 {string} string "Unsupported media type"
// @Router /v1/self/ping [get]
func (h *selfRouter) ping(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	response := pingResponse{
		Response: fmt.Sprint("Pong - ", t.Format(time.RFC3339)),
	}

	h.responseBody(w, r, http.StatusOK, response)
}

// Readiness godoc
// @Summary Respond to an readiness request
// @Description Respond to an readiness request with information about ability of the application to start serving requests.
// @Tags health
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Success 200 {object} info.readinessResponse
// @Router /v1/self/readiness [get]
func (h *selfRouter) readiness(w http.ResponseWriter, r *http.Request) {
	healthStatus, err := h.healthService.Readiness()
	if err != nil {
		healthStatus = &HealthStatus{
			Checks:    make([]HealthCheckResult, 0, 0),
			IsHealthy: false,
		}
	}

	h.readinessResponse(w, r, healthStatus)
}

func (h *selfRouter) readinessResponse(w http.ResponseWriter, r *http.Request, status *HealthStatus) {
	t := time.Now()

	statusText := statusToText(status.IsHealthy)
	responseCode := statusToResponseCode(status.IsHealthy)

	var checkResults []summaryCheckInformation
	checkResults = make([]summaryCheckInformation, 0, len(status.Checks))
	for _, check := range status.Checks {

		result := summaryCheckInformation{
			Name:   check.Name,
			Status: statusToText(check.IsSuccess),
		}
		checkResults = append(checkResults, result)
	}

	response := &readinessResponse{
		Checks:    checkResults,
		Status:    statusText,
		Timestamp: t.Format(time.RFC3339),
	}

	h.responseBody(w, r, responseCode, response)
}

// Started godoc
// @Summary Respond to an started request
// @Description Respond to an started request with information indicating if the application has started successfully.
// @Tags health
// @Accept json
// @Accept xml
// @Produce json
// @Produce xml
// @Success 200 {object} info.startedResponse
// @Router /v1/self/started [get]
func (h *selfRouter) started(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	healthStatus, err := h.healthService.Started()
	if err != nil {
		body := &startedResponse{
			Timestamp: t.Format(time.RFC3339),
		}
		h.responseBody(w, r, http.StatusServiceUnavailable, body)
		return
	}

	body := &startedResponse{
		Timestamp: healthStatus.Timestamp.Format(time.RFC3339),
	}
	h.responseBody(w, r, http.StatusOK, body)
}

func (h *selfRouter) responseBody(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	mediatype, _, err := mime.ParseMediaType(r.Header.Get("Accept"))
	if err != nil {
		observability.LogErrorWithFields(
			log.Fields{
				"error": err},
			"Invalid 'Accept' header.")

		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	switch mediatype {
	case "application/xml":
		render.Status(r, status)
		render.XML(w, r, data)
		return
	case "application/json":
		render.Status(r, status)
		render.JSON(w, r, data)
		return
	default:
		observability.LogErrorWithFields(
			log.Fields{
				"mediatype": mediatype},
			"Invalid media type. Expected either 'application/json' or 'application/xml'")

		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}
}

func statusToResponseCode(status bool) int {
	responseCode := http.StatusOK
	if !status {
		responseCode = http.StatusInternalServerError
	}
	return responseCode
}

func statusToText(status bool) string {
	statusText := Success
	if !status {
		statusText = Failed
	}

	return statusText
}
