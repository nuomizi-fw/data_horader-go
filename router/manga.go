package router

import "github.com/nuomizi-fw/data_horader-go/core"

type MangaRouter struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
}

func NewMangaRouter(data_horader core.DataHoraderServer, logger core.DataHoraderLogger) MangaRouter {
	return MangaRouter{
		data_horader: data_horader,
		logger:       logger,
	}
}

func (mr MangaRouter) InitRouter() {
	mr.logger.Info("Initializing manga router")
}
