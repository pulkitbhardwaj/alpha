package main

import (
	"go.uber.org/fx"

	"github.com/pulkitbhardwaj/alpha/social/auth"
	"github.com/pulkitbhardwaj/alpha/social/internal"
)

func main() {
	fx.New(
		fx.Provide(Config),
		internal.Module,
		auth.Module,
	).
		Run()
}
