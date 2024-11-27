package repo

import (
	"context"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)

func (repo *repository) CreateSong(ctx context.Context, params models.NewSong) (models.Song, error) {
	queryBuilder := sq.Insert("songs").Columns("group_name, song_name, release_date, link, lyrics").
		Values(params.Group, params.Song, params.ReleaseDate, params.Link, params.Lyrics)

	sql, args, err := queryBuilder.Suffix("returning id, group_name, song_name, release_date, link, lyrics").PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return models.Song{}, err
	}

	song := models.Song{}

	err = repo.db.GetConn().QueryRow(ctx, sql, args...).Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Link, &song.Lyrics)

	repo.logger.Debug("Success CREATE song in storage")

	return song, err
}
