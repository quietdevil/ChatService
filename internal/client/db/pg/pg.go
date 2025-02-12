package pg

import (
	"chatservice/internal/client/db"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgPool struct {
	dbc *pgxpool.Pool
}

func NewPgPool(pool *pgxpool.Pool) db.DB {
	return &PgPool{
		dbc: pool,
	}

}

func (p *PgPool) Close() {
	p.dbc.Close()
}

func (p *PgPool) ContextExec(ctx context.Context, q db.Query, a ...any) (pgconn.CommandTag, error) {

	return p.dbc.Exec(ctx, q.QueryStr, a...)
}

func (p *PgPool) ContextQuery(ctx context.Context, q db.Query, a ...any) (pgx.Rows, error) {

	return p.dbc.Query(ctx, q.QueryStr, a...)
}

func (p *PgPool) ContextQueryRow(ctx context.Context, q db.Query, a ...any) pgx.Row {

	return p.dbc.QueryRow(ctx, q.QueryStr, a...)
}

func (p *PgPool) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}
