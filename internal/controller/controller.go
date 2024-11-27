package controller

import (
	"context"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"
	"main/internal/service"
)

type controller struct {
	logger  *slog.Logger
	service service.Service
}

type Controller interface {
	GetSongs(ctx context.Context, params operations.GetSongsParams) ([]*models.Song, error)
	GetLyrics(ctx context.Context, params operations.GetSongsLyricsParams) (string, error)
}

func New(service service.Service, logger *slog.Logger) Controller {
	return &controller{
		logger:  logger,
		service: service,
	}
}
