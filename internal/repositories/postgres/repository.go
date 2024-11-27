//go:generate mockgen -source=./repository.go -destination=../../mocks/repo_mock.go -package=mock
package repo

import (
	"context"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/database"
	"main/internal/models"
)

type repository struct {
	db     database.DB
	logger *slog.Logger
}

type Repository interface {
	GetSongs(ctx context.Context, params operations.GetSongsParams) ([]*models.Song, error)
	GetLyrics(ctx context.Context, params operations.GetSongsLyricsParams) (string, error)
	DeleteSong(ctx context.Context, params operations.DeleteSongsParams) error
	UpdateSong(ctx context.Context, params operations.PatchSongsParams) (models.Song, error)
	CreateSong(ctx context.Context, params models.NewSong) (models.Song, error)
}

func NewUserRepo(db database.DB, logger *slog.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}
