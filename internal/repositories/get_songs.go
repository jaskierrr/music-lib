package repo

import (
	"context"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)

func (repo *repository) GetSongs(ctx context.Context, params operations.GetSongsParams) ([]*models.Song, error) {
	queryBuilder := sq.Select("id, group_name, song_name").From("songs")

	if params.Group != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"group_name": *params.Group})
	}
	if params.Song != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"song_name": *params.Song})
	}
	if params.Link != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"link": *params.Link})
	}
	if params.ReleaseDate != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"release_date": *params.ReleaseDate})
	}

	offset := (*params.Page - 1) * *params.Limit
	queryBuilder = queryBuilder.Limit(uint64(*params.Limit)).Offset(uint64(offset))

	sql, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repo.db.GetConn().Query(ctx, sql, args...)

	songs := []*models.Song{}
	defer rows.Close()

	for rows.Next() {
		song := models.Song{}

		err := rows.Scan(&song.ID, &song.Group, &song.Song)

		if err != nil {
			repo.logger.Error("Error",
				slog.Any("error", err),
			)
			return nil, err
		}

		songs = append(songs, &song)
	}

	repo.logger.Info("Success GET songs from storage")

	return songs, err
}
