package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Hadler func(context.Context) error

type Client interface {
	DB() DB
	Close() error
}

type TxManager interface {
	ReadCommited(context.Context, Hadler) error
}

type Transactor interface {
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
}

type DB interface {
	ExecQuery
	Pinger
	Transactor
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
