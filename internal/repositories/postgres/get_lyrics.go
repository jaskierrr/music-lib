package repo

import (
	"context"
	"main/api/restapi/operations"

	sq "github.com/Masterminds/squirrel"
)


func (repo *repository) GetLyrics(ctx context.Context, params operations.GetSongsLyricsParams) (string, error) {
	queryBuilder := sq.Select("lyrics").From("songs").Where(sq.Eq{"id": params.ID})

	sql, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return "", err
	}

	lyrics := ""

	err = repo.db.GetConn().QueryRow(ctx, sql, args...).Scan(&lyrics)

	repo.logger.Debug("Success GET song lyrics from storage")

	return lyrics, err
}
