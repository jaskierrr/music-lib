package external_repo

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"
	"net/http"
	"net/url"
)

type songAPIClient struct {
	logger     *slog.Logger
	apiURL     string
	httpClient *http.Client
}
type SongAPIClient interface {
	CreateSong(ctx context.Context, params operations.PostSongsParams) (models.NewSong, error)
}

func NewSongAPIClient(apiURL string, logger *slog.Logger) SongAPIClient {
	return &songAPIClient{
		apiURL:     apiURL,
		logger:     logger,
		httpClient: &http.Client{},
	}
}

// http://localhost:8082/info?group=Muse&song=Supermassive Black Hole

func (repo *songAPIClient) CreateSong(ctx context.Context, params operations.PostSongsParams) (models.NewSong, error) {
	queryParams := url.Values{}
	queryParams.Add("group", params.Song.Group)
	queryParams.Add("song", params.Song.Song)

	url := fmt.Sprintf("%s/info?%s", repo.apiURL, queryParams.Encode())

	repo.logger.Debug("Query",
		slog.Any("url", url),
	)

	resp, err := repo.httpClient.Get(url)
	if err != nil {
		return models.NewSong{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.NewSong{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	songDetail := models.SongDetail{}

	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return models.NewSong{}, err
	}

	song := models.NewSong{
		Group:       params.Song.Group,
		Song:        params.Song.Song,
		Link:        songDetail.Link,
		Lyrics:      songDetail.Text,
		ReleaseDate: songDetail.ReleaseDate,
	}

	repo.logger.Debug("Query",
		slog.Any("songDetail", songDetail),
		slog.Any("song", song),
	)

	return song, nil
}
