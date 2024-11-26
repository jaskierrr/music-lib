//go:generate mockgen -source=./service.go -destination=../mocks/service_mock.go -package=mock
package service

import (
	"context"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"
	repo "main/internal/repositories"
)

type service struct {
	logger   *slog.Logger
	repo     repo.Repository
}

type Service interface {
	GetSongs(ctx context.Context, params operations.GetSongsParams) ([]*models.Song, error)
}

func New(repo repo.Repository, logger *slog.Logger) Service {
	return &service{
		logger:   logger,
		repo:     repo,
	}
}
