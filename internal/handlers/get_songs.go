package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetSongs(params operations.GetSongsParams) middleware.Responder {
	h.logger.Debug("Trying to GET songs from storage",
		slog.Any("params", params),
	)

	ctx := params.HTTPRequest.Context()
	songs, err := h.controller.GetSongs(ctx, params)

	if err != nil {
		h.logger.Error(
			"Failed to GET songs from storage",
			slog.String("error", err.Error()),
		)
		return operations.NewGetSongsBadRequest().WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET songs frome storage: " + err.Error(),
			},
		})
	}

	return operations.NewGetSongsOK().WithPayload(songs)
}
