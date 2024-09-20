package service

import "github.com/nuomizi-fw/data-horader/core"

type TorrentService interface{}

type torrentService struct{}

func NewTorrentService(
	db core.DataHoraderDB,
	logger core.DataHoraderLogger,
) TorrentService {
	return &torrentService{}
}
