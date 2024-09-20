package service

import (
	"github.com/nuomizi-fw/data-horader/core"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(NewDataHoraderService),
	),
)

type DataHoraderService struct {
	horader HoraderService
}

func NewDataHoraderService(
	db core.DataHoraderDB,
	logger core.DataHoraderLogger,
) DataHoraderService {
	return DataHoraderService{
		horader: NewHoraderService(db, logger),
	}
}
