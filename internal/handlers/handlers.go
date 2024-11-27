package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/controller"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-playground/validator/v10"
)

type handlers struct {
	logger     *slog.Logger
	controller controller.Controller
}

type Handlers interface {
	GetSongs(params operations.GetSongsParams) middleware.Responder
	GetLyrics(params operations.GetSongsLyricsParams) middleware.Responder

	Link(api *operations.MusicLibraryAPI)
}

func New(controller controller.Controller, validator *validator.Validate, logger *slog.Logger) Handlers {
	return &handlers{
		logger:     logger,
		controller: controller,
	}
}

func (h *handlers) Link(api *operations.MusicLibraryAPI) {
	api.GetSongsHandler = operations.GetSongsHandlerFunc(h.GetSongs)
	api.GetSongsLyricsHandler = operations.GetSongsLyricsHandlerFunc(h.GetLyrics)
}
