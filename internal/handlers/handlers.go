package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/controller"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-playground/validator/v10"
)

type handlers struct {
	logger     *slog.Logger
	controller controller.Controller
}

var validate *validator.Validate

type Handlers interface {
	GetSongs(params operations.GetSongsParams) middleware.Responder

	Link(api *operations.MusicLibraryAPI)
}

func New(controller controller.Controller, validator *validator.Validate, logger *slog.Logger) Handlers {
	validate = validator
	
	return &handlers{
		logger:     logger,
		controller: controller,
	}
}

func (h *handlers) Link(api *operations.MusicLibraryAPI) {
	api.GetSongsHandler = operations.GetSongsHandlerFunc(h.GetSongs)
}

func convertI64tStr(integer int64) string {
	return strconv.FormatInt(integer, 10)
}
