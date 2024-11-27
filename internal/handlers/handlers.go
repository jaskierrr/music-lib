package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/controller"

	"github.com/go-openapi/runtime/middleware"
)

type handlers struct {
	logger     *slog.Logger
	controller controller.Controller
}

type Handlers interface {
	GetSongs(params operations.GetSongsParams) middleware.Responder
	DeleteSong(params operations.DeleteSongsParams) middleware.Responder
	GetLyrics(params operations.GetSongsLyricsParams) middleware.Responder
	UpdateSong(params operations.PatchSongsParams) middleware.Responder
	CreateSong(params operations.PostSongsParams) middleware.Responder

	Link(api *operations.MusicLibraryAPI)
}

func New(controller controller.Controller, logger *slog.Logger) Handlers {
	return &handlers{
		logger:     logger,
		controller: controller,
	}
}

func (h *handlers) Link(api *operations.MusicLibraryAPI) {
	api.GetSongsHandler = operations.GetSongsHandlerFunc(h.GetSongs)
	api.GetSongsLyricsHandler = operations.GetSongsLyricsHandlerFunc(h.GetLyrics)
	api.DeleteSongsHandler = operations.DeleteSongsHandlerFunc(h.DeleteSong)
	api.PatchSongsHandler = operations.PatchSongsHandlerFunc(h.UpdateSong)
	api.PostSongsHandler = operations.PostSongsHandlerFunc(h.CreateSong)
}
