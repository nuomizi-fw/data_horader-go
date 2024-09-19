package router

import "github.com/nuomizi-fw/data_horader-go/core"

type RssRouter struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
}

func NewRssRouter(data_horader core.DataHoraderServer, logger core.DataHoraderLogger) RssRouter {
	return RssRouter{
		data_horader: data_horader,
		logger:       logger,
	}
}

func (rr RssRouter) InitRouter() {
	rr.logger.Info("Initializing Rss router")
}
