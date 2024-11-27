package repo

import (
	"context"
	"errors"
	"log/slog"
	"main/api/restapi/operations"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)

func (repo *repository) GetSongs(ctx context.Context, params operations.GetSongsParams) ([]*models.Song, error) {
	queryBuilder := sq.Select("id, group_name, song_name, link, release_date, lyrics").From("songs")

	if params.Song.ID != 0 {
		queryBuilder = queryBuilder.Where(sq.Eq{"id": params.Song.ID})
	}
	if params.Song.Group != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"group_name": params.Song.Group})
	}
	if params.Song.Song != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"song_name": params.Song.Song})
	}
	if params.Song.Link != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"link": params.Song.Link})
	}
	if params.Song.ReleaseDate != "" {
		queryBuilder = queryBuilder.Where(sq.Eq{"release_date": params.Song.ReleaseDate})
	}

	offset := (params.Page - 1) * params.Limit
	queryBuilder = queryBuilder.Limit(uint64(params.Limit)).Offset(uint64(offset))

	sql, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	repo.logger.Debug("Query",
		slog.Any("sql", sql),
		slog.Any("args", args),
	)


	rows, err := repo.db.GetConn().Query(ctx, sql, args...)
	rowsCount := 0
	songs := []*models.Song{}
	defer rows.Close()

	for rows.Next() {
		rowsCount++
		song := models.Song{}

		err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.Link, &song.ReleaseDate, &song.Lyrics)

		if err != nil {
			repo.logger.Error("Error",
				slog.Any("error", err),
			)
			return nil, err
		}


		repo.logger.Debug("WTF",
			slog.Any("row", song),
		)
		songs = append(songs, &song)
	}

	if rowsCount == 0 {
		repo.logger.Error("no songs found")
		return nil, errors.New("no songs found")
	}

	repo.logger.Debug("Success GET songs from storage")

	return songs, err
}
