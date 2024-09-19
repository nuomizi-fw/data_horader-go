package router

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"router",
	fx.Provide(
		NewDataHoraderRouter,
	),
	// Add new router below
	fx.Provide(
		NewAuthRouter,
		NewUserRouter,
		NewAria2Router,
		NewBittorrentRouter,
		NewMangaRouter,
		NewMusicRouter,
		NewNovelRouter,
		NewSearchRouter,
		NewBangumiRouter,
	),
)

type DataHoraderRouter interface {
	InitRouter()
}

type DataHoraderRouters []DataHoraderRouter

func (sr DataHoraderRouters) InitRouter() {
	for _, router := range sr {
		router.InitRouter()
	}
}

func NewDataHoraderRouter(
	authRouter AuthRouter,
	userRouter UserRouter,
	aria2Router Aria2Router,
	bittorrentRouter BittorrentRouter,
	mangaRouter MangaRouter,
	musicRouter MusicRouter,
	novelRouter NovelRouter,
	searchRouter SearchRouter,
	bangumiRouter BangumiRouter,
) DataHoraderRouters {
	return DataHoraderRouters{
		authRouter,
		userRouter,
		aria2Router,
		bittorrentRouter,
		mangaRouter,
		musicRouter,
		novelRouter,
		searchRouter,
		bangumiRouter,
	}
}
