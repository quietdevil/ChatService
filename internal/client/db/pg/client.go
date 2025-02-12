package pg

import (
	"chatservice/internal/client/db"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBClient struct {
	db db.DB
}

func NewDBClient(ctx context.Context, dsn string) (db.Client, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &DBClient{
		db: &PgPool{dbc: pool},
	}, nil

}

func (c *DBClient) Close() error {
	c.db.Close()
	return nil
}

func (c *DBClient) DB() db.DB {
	return c.db
}
