package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Client interface {
	DB() DB
	Close() error
}

type DB interface {
	ExecQuery
	Pinger
	Close()
}

type Query struct {
	Name     string
	QueryStr string
}

type ExecQuery interface {
	ContextExec(context.Context, Query, ...any) (pgconn.CommandTag, error)
	ContextQuery(context.Context, Query, ...any) (pgx.Rows, error)
	ContextQueryRow(context.Context, Query, ...any) pgx.Row
}

type Pinger interface {
	Ping(context.Context) error
}
