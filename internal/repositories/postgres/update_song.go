package repo

import (
	"context"
	"main/api/restapi/operations"
	"main/internal/models"

	sq "github.com/Masterminds/squirrel"
)

func (repo *repository) UpdateSong(ctx context.Context, params operations.PatchSongsParams) (models.Song, error) {
	queryBuilder := sq.Update("songs").Where(sq.Eq{"id": params.Song.ID})

	if params.Song.Group != "" {
		queryBuilder = queryBuilder.Set("group_name", params.Song.Group)
	}
	if params.Song.Song != "" {
		queryBuilder = queryBuilder.Set("song_name", params.Song.Song)
	}
	if params.Song.Link != "" {
		queryBuilder = queryBuilder.Set("link", params.Song.Link)
	}
	if params.Song.ReleaseDate != "" {
		queryBuilder = queryBuilder.Set("release_date", params.Song.ReleaseDate)
	}
	if params.Song.Lyrics != "" {
		queryBuilder = queryBuilder.Set("lyrics", params.Song.Lyrics)
	}


	sql, args, err := queryBuilder.Suffix("returning id, group_name, song_name, release_date, link, lyrics").PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return models.Song{}, err
	}

	song := models.Song{}

	err = repo.db.GetConn().QueryRow(ctx, sql, args...).Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Link, &song.Lyrics)

	repo.logger.Debug("Success UPDATE song in storage")

	return song, err
}
