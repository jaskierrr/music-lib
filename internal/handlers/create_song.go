package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) CreateSong(params operations.PostSongsParams) middleware.Responder {
	h.logger.Debug("Trying to CREATE song in storage",
		slog.Any("params", params),
	)

	ctx := params.HTTPRequest.Context()
	song, err := h.controller.CreateSong(ctx, params)

	if err != nil {
		h.logger.Error(
			"Failed to CREATE song in storage",
			slog.String("error", err.Error()),
		)
		return operations.NewPostSongsBadRequest().WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to CREATE song in storage: " + err.Error(),
			},
		})
	}

	return operations.NewPostSongsCreated().WithPayload(&song)
}
