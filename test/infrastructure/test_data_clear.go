package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

func TruncateTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	query := "TRUNCATE books RESTART IDENTITY"

	_, err := dbPool.Exec(ctx, query)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Books table truncated")
	}
}
