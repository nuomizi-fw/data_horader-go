package router

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"router",
	fx.Provide(
		NewDataHoraderRouter,
	),
	// Add new router below
	fx.Provide(
		NewTorrentRouter,
	),
)

type DataHoraderRouter interface {
	InitRouter()
}

type DataHoraderRouters []DataHoraderRouter

func (sr DataHoraderRouters) InitRouter() {
	for _, router := range sr {
		router.InitRouter()
	}
}

func NewDataHoraderRouter(
	torrentRouter TorrentRouter,
) DataHoraderRouters {
	return DataHoraderRouters{
		torrentRouter,
	}
}
