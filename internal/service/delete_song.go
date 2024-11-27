package service

import (
	"context"
	"main/api/restapi/operations"
)

func (s service) DeleteSong(ctx context.Context, params operations.DeleteSongsParams) error {
	err := s.repo.DeleteSong(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
