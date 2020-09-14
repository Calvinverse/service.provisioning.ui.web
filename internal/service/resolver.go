package service

import (
	"github.com/spf13/cobra"

	"github.com/calvinverse/service.provisioning/internal/cmd"
)

// Resolver defines the interface for Inversion-of-Control objects.
type Resolver interface {
	ResolveCommands() []*cobra.Command
}

// NewResolver returns a new Resolver instance
func NewResolver(config Configuration) Resolver {
	return &concreteResolver{
		config: config,
	}
}

// Resolver provides the data and methods to resolve types in the service.
type concreteResolver struct {
	config Configuration

	commands []*cobra.Command
}

func (r *concreteResolver) ResolveRouter() *foo.HTTPClient {
	return foo.NewHTTPClient(
		r.ResolveLogger(),
		r.config.HTTP,
	)
}

// ResolveCommands returns a collection of commands for the application.
func (r *concreteResolver) ResolveCommands() []*cobra.Command {
	if r.commands == nil {
		r.commands = []*cobra.Command{
			cmd.ServerCmd,
		}
	}

	return r.commands
}
