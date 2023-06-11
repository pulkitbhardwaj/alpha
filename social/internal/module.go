package internal

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewExecutableSchema),
	fx.Provide(SQLGraph),
	fx.Invoke(HTTPServer),
)
