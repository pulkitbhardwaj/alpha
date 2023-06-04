package main

import "go.uber.org/fx"

func main() {
	fx.New(
		fx.Provide(NewSQLGraph),
		fx.Provide(NewGraphQLSchema),
		fx.Provide(NewGraphQLHandler),
		fx.Invoke(NewHTTPServer),
	).Run()
}
