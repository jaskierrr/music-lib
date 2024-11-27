package controller

import (
	"context"
	"main/api/restapi/operations"
)

func (c controller) GetLyrics(ctx context.Context, params operations.GetSongsLyricsParams) (string, error) {
	return c.service.GetLyrics(ctx, params)
}
