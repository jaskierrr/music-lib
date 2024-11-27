//go:generate mockgen -source=./service.go -destination=../mocks/service_mock.go -package=mock
package service

import (
	"context"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"
	external_repo "main/internal/repositories/externalApi"
	repo "main/internal/repositories/postgres"
)

type service struct {
	logger *slog.Logger
	repo   repo.Repository
	externalRepo external_repo.SongAPIClient
}

type Service interface {
	GetSongs(ctx context.Context, params operations.GetSongsParams) ([]*models.Song, error)
	GetLyrics(ctx context.Context, params operations.GetSongsLyricsParams) (string, error)
	DeleteSong(ctx context.Context, params operations.DeleteSongsParams) error
	UpdateSong(ctx context.Context, params operations.PatchSongsParams) (models.Song, error)
	CreateSong(ctx context.Context, params operations.PostSongsParams) (models.Song, error)
}

func New(repo repo.Repository, externalRepo external_repo.SongAPIClient, logger *slog.Logger) Service {
	return &service{
		repo:   repo,
		externalRepo: externalRepo,
		logger: logger,
	}
}
