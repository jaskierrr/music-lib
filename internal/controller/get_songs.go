package controller

import (
	"context"
	"main/api/restapi/operations"
	"main/internal/models"
)

func (c controller) GetSongs(ctx context.Context, params operations.GetSongsParams) ([]*models.Song, error) {
	return c.service.GetSongs(ctx, params)
}
