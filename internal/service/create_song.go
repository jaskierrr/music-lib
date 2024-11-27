package service

import (
	"context"
	"main/api/restapi/operations"
	"main/internal/models"
)

func (s service) CreateSong(ctx context.Context, params operations.PostSongsParams) (models.Song, error) {
	songDetail, err := s.externalRepo.CreateSong(ctx, params)
	if err != nil {
		return models.Song{}, err
	}

	song, err := s.repo.CreateSong(ctx, songDetail)
	if err != nil {
		return models.Song{}, err
	}
	
	return song, nil
}
