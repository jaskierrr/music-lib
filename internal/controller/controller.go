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
	DeleteSong(ctx context.Context, params operations.DeleteSongsParams) error
	UpdateSong(ctx context.Context, params operations.PatchSongsParams) (models.Song, error)
	CreateSong(ctx context.Context, params operations.PostSongsParams) (models.Song, error)
}

func New(service service.Service, logger *slog.Logger) Controller {
	return &controller{
		logger:  logger,
		service: service,
	}
}

func (c controller) GetSongs(ctx context.Context, params operations.GetSongsParams) ([]*models.Song, error) {
	return c.service.GetSongs(ctx, params)
}

func (c controller) GetLyrics(ctx context.Context, params operations.GetSongsLyricsParams) (string, error) {
	return c.service.GetLyrics(ctx, params)
}
func (c controller) DeleteSong(ctx context.Context, params operations.DeleteSongsParams) error {
	return c.service.DeleteSong(ctx, params)
}
func (c controller) UpdateSong(ctx context.Context, params operations.PatchSongsParams) (models.Song, error) {
	return c.service.UpdateSong(ctx, params)
}
func (c controller) CreateSong(ctx context.Context, params operations.PostSongsParams) (models.Song, error) {
	return c.service.CreateSong(ctx, params)
}
