package service

import "github.com/nuomizi-fw/data-horader/core"

type HoraderService interface{}

type horaderService struct{}

func NewHoraderService(
	db core.DataHoraderDB,
	logger core.DataHoraderLogger,
) HoraderService {
	return &horaderService{}
}
