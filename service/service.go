package service

import (
	"github.com/nuomizi-fw/data-horader/core"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(NewDataHoraderService),
		// Add new service below
		fx.Provide(
			NewTorrentService,
		),
	),
)

type DataHoraderService struct {
	torrent TorrentService
}

func NewDataHoraderService(
	db core.DataHoraderDB,
	logger core.DataHoraderLogger,
) DataHoraderService {
	return DataHoraderService{
		torrent: NewTorrentService(db, logger),
	}
}
