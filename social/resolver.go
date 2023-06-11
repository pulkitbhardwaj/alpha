package main

import (
	"go.uber.org/fx"

	"github.com/pulkitbhardwaj/alpha/social/internal"
)

// This file will not be regenerated automatically.

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	fx.In

	// graph internal.Graph
}

func Config(r Resolver) internal.Config {
	return internal.Config{
		Resolvers: &r,
	}
}
