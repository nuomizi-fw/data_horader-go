package router

import "github.com/nuomizi-fw/data_horader-go/core"

type SearchRouter struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
}

func NewSearchRouter(data_horader core.DataHoraderServer, logger core.DataHoraderLogger) SearchRouter {
	return SearchRouter{
		data_horader: data_horader,
		logger:       logger,
	}
}

func (sr SearchRouter) InitRouter() {
	sr.logger.Info("Initializing search router")
}
