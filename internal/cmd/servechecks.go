package cmd

import (
	"time"

	"github.com/AppsFlyer/go-sundheit/checks"
)

// ServeLivelinessCheck returns a liveliness check for the serve command.
func ServeLivelinessCheck() checks.Check {
	t := time.NewTicker(5 * time.Second)
	check := &serveLivelinessCheck{}
	check.configureTicker(t)

	return check
}

type serveLivelinessCheck struct {
	lastError error

	ticker *time.Ticker
}

func (s *serveLivelinessCheck) Execute() (details interface{}, err error) {
	details = "check-details"

	return details, nil
}

func (s *serveLivelinessCheck) Name() string {
	return "serve-liveliness"
}

func (s *serveLivelinessCheck) configureTicker(ticker *time.Ticker) {
	s.ticker = ticker
	go func() {
		for {
			select {
			case <-ticker.C:
				s.lastError = nil
			}
		}
	}()
}
