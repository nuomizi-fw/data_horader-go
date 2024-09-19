package middleware

import "go.uber.org/fx"

var Module = fx.Module(
	"middleware",
	fx.Options(
		fx.Provide(NewMiddleware),
		// Add new middleware below
		fx.Provide(NewJWTMiddleware),
		fx.Provide(NewCorsMiddleware),
	),
)

type DataHoraderMiddleware interface {
	InitMiddleware()
}

type DataHoraderMiddlewares []DataHoraderMiddleware

func (sm DataHoraderMiddlewares) InitMiddleware() {
	for _, middleware := range sm {
		middleware.InitMiddleware()
	}
}

func NewMiddleware() DataHoraderMiddlewares {
	return DataHoraderMiddlewares{}
}
