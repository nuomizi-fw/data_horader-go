package model

import (
	"github.com/nuomizi-fw/data_horader-go/internal/core"
	"go.uber.org/fx"
)

var Module = fx.Module("model", fx.Provide(NewDataHoraderModels))

type DataHoraderModel struct {
	models []interface{}
	db     core.DataHoraderDB
	logger core.DataHoraderLogger
	config core.DataHoraderConfig
}

func (sm DataHoraderModel) AutoMigrate() {
	if !sm.config.Database.Migrate {
		sm.logger.Info("Database migration is disabled")
		return
	}

	sm.logger.Info("Auto migrating models")
	if err := sm.db.AutoMigrate(sm.models...); err != nil {
		sm.logger.Error("Failed to auto migrate models", err)
	}
}

func NewDataHoraderModels(
	db core.DataHoraderDB,
	logger core.DataHoraderLogger,
	config core.DataHoraderConfig,
) DataHoraderModel {
	return DataHoraderModel{
		db:     db,
		logger: logger,
		config: config,
		models: []interface{}{
			&User{},
			&Bangumi{},
			&Season{},
			&CastMember{},
		},
	}
}
