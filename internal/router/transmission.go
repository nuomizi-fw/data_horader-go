package router

import "github.com/nuomizi-fw/data_horader-go/internal/core"

type TransmissionRouter struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
}

func NewTransmissionRouter(data_horader core.DataHoraderServer, logger core.DataHoraderLogger) TransmissionRouter {
	return TransmissionRouter{
		data_horader: data_horader,
		logger:       logger,
	}
}

func (tr TransmissionRouter) InitRouter() {
	tr.logger.Info("Initializing transmission router")
}
