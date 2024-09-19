package core

import "go.uber.org/fx"

var Module = fx.Module(
	"core",
	fx.Options(
		fx.Provide(
			NewDataHoraderConfig,
			NewDataHoraderDB,
			NewDataHoraderLogger,
			NewDataHoraderServer,
		),
	),
)
