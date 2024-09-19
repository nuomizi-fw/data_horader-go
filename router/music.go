package router

import "github.com/nuomizi-fw/data_horader-go/core"

type MusicRouter struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
}

func NewMusicRouter(data_horader core.DataHoraderServer, logger core.DataHoraderLogger) MusicRouter {
	return MusicRouter{
		data_horader: data_horader,
		logger:       logger,
	}
}

func (mr MusicRouter) InitRouter() {
	mr.logger.Info("Initializing music router")
}
