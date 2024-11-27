package handlers

import (
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	"github.com/go-openapi/runtime/middleware"
)

func (h *handlers) GetLyrics(params operations.GetSongsLyricsParams) middleware.Responder {
	h.logger.Debug("Trying to GET song lyrics from storage",
		slog.Any("params", params),
	)

	ctx := params.HTTPRequest.Context()
	lyrics, err := h.controller.GetLyrics(ctx, params)

	if err != nil {
		h.logger.Error(
			"Failed to GET song lyrics from storage",
			slog.String("error", err.Error()),
		)
		return operations.NewGetSongsLyricsBadRequest().WithPayload(&models.ErrorResponse{
			Error: &models.ErrorResponseAO0Error{
				Message: "Failed to GET song lyrics frome storage: " + err.Error(),
			},
		})
	}

	return operations.NewGetSongsLyricsOK().WithPayload(lyrics)
}
