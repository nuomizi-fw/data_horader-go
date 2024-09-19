package router

import "github.com/nuomizi-fw/data_horader-go/core"

type NovelRouter struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
}

func NewNovelRouter(data_horader core.DataHoraderServer, logger core.DataHoraderLogger) NovelRouter {
	return NovelRouter{
		data_horader: data_horader,
		logger:       logger,
	}
}

func (nr NovelRouter) InitRouter() {
	nr.logger.Info("Initializing novel router")
}
