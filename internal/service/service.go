package service

import (
	"github.com/nuomizi-fw/data_horader-go/internal/core"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(NewDataHoraderService),
		// Add new service below
		fx.Provide(
			NewAuthService,
			NewUserService,
			NewAria2Service,
			NewBittorrentService,
			NewTransmissionService,
			NewMangaService,
			NewMusicService,
			NewNovelService,
			NewSearchService,
			NewBangumiService,
		),
	),
)

type DataHoraderService struct {
	auth         AuthService
	user         UserService
	rss          RssService
	aria2        Aria2Service
	bittorrent   BittorrentService
	transmission TransmissionService
	manga        MangaService
	music        MusicService
	search       SearchService
	bangumi      BangumiService
}

func NewDataHoraderService(
	db core.DataHoraderDB,
	logger core.DataHoraderLogger,
) DataHoraderService {
	return DataHoraderService{
		auth:       NewAuthService(db, logger),
		user:       NewUserService(db, logger),
		aria2:      NewAria2Service(),
		bittorrent: NewBittorrentService(),
		manga:      NewMangaService(),
		music:      NewMusicService(),
		search:     NewSearchService(),
		bangumi:    NewBangumiService(),
	}
}
