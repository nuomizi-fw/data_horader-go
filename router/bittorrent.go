package router

import (
	"github.com/nuomizi-fw/data_horader-go/core"
)

type BittorrentRouter struct {
	data_horader core.DataHoraderServer
	logger       core.DataHoraderLogger
}

func NewBittorrentRouter(data_horader core.DataHoraderServer, logger core.DataHoraderLogger) BittorrentRouter {
	return BittorrentRouter{
		data_horader: data_horader,
		logger:       logger,
	}
}

func (br BittorrentRouter) InitRouter() {
	br.logger.Info("Initializing bittorrent router")
}
