package core

import (
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DataHoraderDB struct {
	*gorm.DB
}

func NewDataHoraderDB(config DataHoraderConfig, logger DataHoraderLogger) DataHoraderDB {
	db, err := gorm.Open(sqlite.Open(config.Database.DBFile), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})
	if err != nil {
		logger.Error("Failed to connect to database", err)
	}

	c, _ := db.DB()
	c.SetMaxIdleConns(10)
	c.SetMaxOpenConns(1)
	c.SetConnMaxIdleTime(time.Second * 1000)

	return DataHoraderDB{db}
}
