package service

import (
	"context"
	"main/api/restapi/operations"
	"main/internal/models"
)

func (s service) UpdateSong(ctx context.Context, params operations.PatchSongsParams) (models.Song, error) {
	song, err := s.repo.UpdateSong(ctx, params)
	if err != nil {
		return models.Song{}, err
	}

	return song, nil
}
