package repo

import (
	"context"
	"main/api/restapi/operations"

	sq "github.com/Masterminds/squirrel"
)

func (repo *repository) DeleteSong(ctx context.Context, params operations.DeleteSongsParams) error {
	queryBuilder := sq.Delete("songs").Where(sq.Eq{"id": params.ID})

	sql, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = repo.db.GetConn().Exec(ctx, sql, args...)

	repo.logger.Debug("Success DELETE song from storage")

	return err
}
