package service

import (
	"context"
	"main/api/restapi/operations"
	"main/internal/models"
)

func (s service) GetSongs(ctx context.Context, params operations.GetSongsParams) ([]*models.Song, error) {
	users, err := s.repo.GetSongs(ctx, params)
	if err != nil {
		return nil, err
	}

	return users, nil
}
