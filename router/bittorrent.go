package router

import (
	"github.com/nuomizi-fw/data-horader/core"
)

type TorrentRouter struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
}

func NewTorrentRouter(data_horader core.DataHoraderServer, logger core.DataHoraderLogger) TorrentRouter {
	return TorrentRouter{
		data_horader: data_horader,
		logger:       logger,
	}
}

func (tr TorrentRouter) InitRouter() {
	tr.logger.Info("Initializing torrent router")
}
