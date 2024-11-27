package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) UpdateSong(params operations.PatchSongsParams) middleware.Responder {
	h.logger.Debug("Trying to UPDATE song in storage",
		slog.Any("params", params),
	)

	ctx := params.HTTPRequest.Context()
	song, err := h.controller.UpdateSong(ctx, params)

	if err != nil {
		h.logger.Error(
			"Failed to UPDATE song in storage",
			slog.String("error", err.Error()),
		)
		return operations.NewPatchSongsBadRequest().WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to UPDATE song in storage: " + err.Error(),
			},
		})
	}

	return operations.NewPatchSongsOK().WithPayload(&song)
}
