package info

import (
	"testing"
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"
)

//
// Mocks
//

type mockCheck struct {
	CheckName string
	Counter   int
}

func (m *mockCheck) Name() string {
	return m.CheckName
}

func (m *mockCheck) Execute() (details interface{}, err error) {
	m.Counter++
	return m.Counter, nil
}

func createMockHealthInstance(healthy bool) *mockHealthInstance {
	return &mockHealthInstance{
		checks:    make([]*gosundheit.Config, 0),
		isHealthy: healthy,
	}
}

type mockHealthInstance struct {
	checks    []*gosundheit.Config
	isHealthy bool
}

func (m *mockHealthInstance) RegisterCheck(cfg *gosundheit.Config) error {
	m.checks = append(m.checks, cfg)

	return nil
}

func (m *mockHealthInstance) Deregister(name string) {
	// do nothing
}

func (m *mockHealthInstance) Results() (results map[string]gosundheit.Result, healthy bool) {
	var resultMap map[string]gosundheit.Result
	resultMap = make(map[string]gosundheit.Result)
	for _, v := range m.checks {
		startTime := time.Now()

		details, err := v.Check.Execute()

		duration := time.Since(startTime)

		result := gosundheit.Result{
			Details:            details,
			Error:              err,
			Timestamp:          startTime,
			Duration:           duration,
			TimeOfFirstFailure: nil,
		}

		resultMap[v.Check.Name()] = result
	}

	return resultMap, m.isHealthy
}

func (m *mockHealthInstance) IsHealthy() bool {
	return m.isHealthy
}

func (m *mockHealthInstance) DeregisterAll() {
	// do nothing
}

func (m *mockHealthInstance) WithCheckListener(listener gosundheit.CheckListener) {
	// do nothing
}

//
// StatusReporter
//

// liveliness
func TestLivelinessWithOneCheck(t *testing.T) {
	mockCheck := &mockCheck{
		CheckName: "a",
	}

	mock := createMockHealthInstance(true)
	setHealthInstanceForTesting(mock)

	healthInstance := GetHealthCenter()
	healthInstance.RegisterLivelinessCheck(mockCheck, 10*time.Second, 10*time.Second, true)

	reporter := GetStatusReporter()

	status, err := reporter.Liveliness()
	if err != nil {
		t.Errorf("Getting status reporter failed with error: %s", err.Error())
	}

	checks := status.Checks
	if len(checks) != 1 {
		t.Errorf("Got %d checks, expected 1 checks", len(status.Checks))
	}

	if checks[0].Name != mockCheck.Name() {
		t.Errorf("Got check with name %s, expected %s", checks[0].Name, mockCheck.Name())
	}

	if !checks[0].IsSuccess {
		t.Errorf("Check with name %s was not successful", checks[0].Name)
	}
}

func TestLivelinessWithTwoChecks(t *testing.T) {
	mockCheck1 := &mockCheck{
		CheckName: "a",
	}

	mockCheck2 := &mockCheck{
		CheckName: "b",
	}

	mock := createMockHealthInstance(true)
	setHealthInstanceForTesting(mock)

	healthInstance := GetHealthCenter()
	healthInstance.RegisterLivelinessCheck(mockCheck1, 10*time.Second, 10*time.Second, true)
	healthInstance.RegisterLivelinessCheck(mockCheck2, 10*time.Second, 10*time.Second, true)

	reporter := GetStatusReporter()

	status, err := reporter.Liveliness()
	if err != nil {
		t.Errorf("Getting status reporter failed with error: %s", err.Error())
	}

	checks := status.Checks
	if len(checks) != 2 {
		t.Errorf("Got %d checks, expected 2 checks", len(status.Checks))
	}

	if checks[0].Name != mockCheck1.Name() && checks[0].Name != mockCheck2.Name() {
		t.Errorf("Got check with name %s, expected %s or %s", checks[0].Name, mockCheck1.Name(), mockCheck2.Name())
	}

	if !checks[0].IsSuccess {
		t.Errorf("Check with name %s was not successful", checks[0].Name)
	}

	if checks[1].Name != mockCheck1.Name() && checks[1].Name != mockCheck2.Name() {
		t.Errorf("Got check with name %s, expected %s or %s", checks[1].Name, mockCheck1.Name(), mockCheck2.Name())
	}

	if !checks[1].IsSuccess {
		t.Errorf("Check with name %s was not successful", checks[1].Name)
	}
}

func TestLivelinessWithZeroChecks(t *testing.T) {
	mock := createMockHealthInstance(true)
	setHealthInstanceForTesting(mock)

	reporter := GetStatusReporter()

	status, err := reporter.Liveliness()
	if err != nil {
		t.Errorf("Getting status reporter failed with error: %s", err.Error())
	}

	if len(status.Checks) != 0 {
		t.Errorf("Got %d checks, expected 0 checks", len(status.Checks))
	}
}

// readiness

//
// HealthCenter
//
