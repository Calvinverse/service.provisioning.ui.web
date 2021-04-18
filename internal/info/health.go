package info

import (
	"sync"
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/AppsFlyer/go-sundheit/checks"
	"github.com/calvinverse/service.provisioning.ui.web/internal/observability"

	log "github.com/sirupsen/logrus"
)

const (
	// Failed indicates the health or a health check is failing.
	Failed string = "failed"

	// Success indicates the health or a health check is successful.
	Success string = "success"
)

var (
	once     sync.Once
	instance *healthReporter
)

// StatusReporter defines a service that tracks the health of the application.
type StatusReporter interface {
	// Liveliness returns the status indicating if the application is healthy while processing requests.
	Liveliness() (*HealthStatus, error)

	// Readiness returns the status indicating if the application is ready to process requests.
	Readiness() (*HealthStatus, error)

	// Started returns the information about the start of the application.
	Started() (*StartedStatus, error)
}

// HealthCenter defines a service which allows different health checks to be registered for monitoring.
type HealthCenter interface {
	// Add a health check
	RegisterLivelinessCheck(check checks.Check, executionPeriod time.Duration, initialDelay time.Duration, initiallyPassing bool) error

	// Add a readiness check
	// Add a started check
}

// HealthStatus stores the health status for the application.
type HealthStatus struct {
	// Checks returns the collection of checks that were executed.
	Checks []HealthCheckResult

	// IsHealthy returns the health status for the application.
	IsHealthy bool
}

// HealthCheckResult stores the results of a health check.
type HealthCheckResult struct {
	// Description returns the description of the check status.
	Description string

	// IsSuccess returns the status of the check.
	IsSuccess bool

	// Name returns the name of the check.
	Name string

	// The last time the check result was updated.
	Timestamp time.Time
}

// StartedStatus stores the application start information.
type StartedStatus struct {
	// The time the application was started.
	Timestamp time.Time
}

// GetStatusReporter returns a health status reporter which reports on the status of the application.
func GetStatusReporter() StatusReporter {
	once.Do(func() {
		if instance == nil {
			instance = &healthReporter{
				instance: gosundheit.New(),
			}
		}
	})

	return instance
}

// GetHealthCenter returns a HealthCenter instance that can be used to register health checks.
func GetHealthCenter() HealthCenter {
	once.Do(func() {
		if instance == nil {
			instance = &healthReporter{
				instance: gosundheit.New(),
			}
		}
	})

	return instance
}

// setHealthInstanceForTesting sets a status reporter with the provided health instance for testing purposes.
func setHealthInstanceForTesting(healthInstance gosundheit.Health) {
	instance = &healthReporter{
		instance: healthInstance,
	}
}

type healthReporter struct {
	instance gosundheit.Health
}

func (h *healthReporter) Liveliness() (*HealthStatus, error) {
	checkResults, healthy := h.instance.Results()

	var checks []HealthCheckResult
	checks = make([]HealthCheckResult, 0, len(checkResults))
	for name, check := range checkResults {
		checkResult := HealthCheckResult{
			Description: check.String(),
			IsSuccess:   check.IsHealthy(),
			Name:        name,
			Timestamp:   check.Timestamp,
		}

		checks = append(checks, checkResult)
	}

	// Return the status of the different health checks
	result := &HealthStatus{
		Checks:    checks,
		IsHealthy: healthy,
	}
	return result, nil
}

func (h *healthReporter) Readiness() (*HealthStatus, error) {
	// If all health checks have been registered then we are good
	return &HealthStatus{}, nil
}

func (h *healthReporter) RegisterLivelinessCheck(check checks.Check, executionPeriod time.Duration, initialDelay time.Duration, initiallyPassing bool) error {
	err := h.instance.RegisterCheck(&gosundheit.Config{
		Check:            check,
		ExecutionPeriod:  executionPeriod,
		InitialDelay:     initialDelay,
		InitiallyPassing: initiallyPassing,
	})

	if err != nil {
		observability.LogErrorWithFields(
			log.Fields{
				"check_name": check.Name(),
				"error":      err},
			"Failed to register a liveliness check")

		return err
	}

	return nil
}

func (h *healthReporter) Started() (*StartedStatus, error) {
	return &StartedStatus{}, nil
}
