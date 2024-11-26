package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

// Request URL:
// http://localhost:8080/songs?group=Muse&song=Supermassive%20Black%20Hole&link=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DXsp3_a-PMTw&releaseDate=2006-07-16&page=1&limit=3


func (h *handlers) GetSongs(params operations.GetSongsParams) middleware.Responder {
	h.logger.Info("Trying to GET users from storage")

	ctx := params.HTTPRequest.Context()
	users, err := h.controller.GetSongs(ctx, params)

	if err != nil {
		h.logger.Error(
			"Failed to GET users from storage",
			slog.String("error", err.Error()),
		)
		return operations.NewGetSongsBadRequest().WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET songs frome storage " + err.Error(),
			},
		})
	}

	return operations.NewGetSongsOK().WithPayload(users)
}
