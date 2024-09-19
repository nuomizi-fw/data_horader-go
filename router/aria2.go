package router

import (
	"github.com/nuomizi-fw/data_horader-go/core"
)

type Aria2Router struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
}

func NewAria2Router(data_horader core.DataHoraderServer, logger core.DataHoraderLogger) Aria2Router {
	return Aria2Router{
		data_horader: data_horader,
		logger:       logger,
	}
}

func (ar Aria2Router) InitRouter() {
	ar.logger.Info("Initializing aria2 router")
}
