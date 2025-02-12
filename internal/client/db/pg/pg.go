package pg

import (
	"chatservice/internal/client/db"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type key string

const TxKey key = "TxKey"

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

	if tx, ok := ctx.Value(TxKey).(pgx.Tx); ok {
		return tx.Exec(ctx, q.QueryStr, a...)
	}

	return p.dbc.Exec(ctx, q.QueryStr, a...)
}

func (p *PgPool) ContextQuery(ctx context.Context, q db.Query, a ...any) (pgx.Rows, error) {
	if tx, ok := ctx.Value(TxKey).(pgx.Tx); ok {
		return tx.Query(ctx, q.QueryStr, a...)
	}
	return p.dbc.Query(ctx, q.QueryStr, a...)
}

func (p *PgPool) ContextQueryRow(ctx context.Context, q db.Query, a ...any) pgx.Row {

	if tx, ok := ctx.Value(TxKey).(pgx.Tx); ok {
		return tx.QueryRow(ctx, q.QueryStr, a...)
	}

	return p.dbc.QueryRow(ctx, q.QueryStr, a...)
}

func (p *PgPool) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

func (p *PgPool) BeginTx(ctx context.Context, opt pgx.TxOptions) (pgx.Tx, error) {

	return p.dbc.BeginTx(ctx, opt)
}

func MakeContext(ctx context.Context, tx pgx.Tx) context.Context {
	newCtx := context.WithValue(ctx, TxKey, tx)
	return newCtx
}
