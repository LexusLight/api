package api

import (
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	Beatmaps    *echo.Group
	BeatmapSets *echo.Group
	Me          *echo.Group
}

func New(store store.Store, prefix *echo.Group) {
	api := Routes{}

	api.Beatmaps = prefix.Group("/beatmaps")
	api.BeatmapSets = prefix.Group("/beatmapsets")
	api.Me = prefix.Group("/me")

	api.InitBeatmaps(store)
	api.InitBeatmapSet(store)
	api.InitMe(store)
}
