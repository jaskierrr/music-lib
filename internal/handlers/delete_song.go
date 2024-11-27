package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) DeleteSong(params operations.DeleteSongsParams) middleware.Responder {
	h.logger.Debug("Trying to DELETE song from storage",
		slog.Any("params", params),
	)

	ctx := params.HTTPRequest.Context()
	err := h.controller.DeleteSong(ctx, params)

	if err != nil {
		h.logger.Error(
			"Failed to DELETE song from storage",
			slog.String("error", err.Error()),
		)
		return operations.NewDeleteSongsBadRequest().WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to DELETE song frome storage: " + err.Error(),
			},
		})
	}

	return operations.NewDeleteSongsNoContent()
}
