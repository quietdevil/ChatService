package logs

import (
	"chatservice/internal/client/db"
	"chatservice/internal/repository"
	"context"
)

type logs struct {
	dbClient db.Client
}

func NewLogger(dbClient db.Client) repository.Logger {
	return &logs{
		dbClient: dbClient,
	}
}

func (l *logs) Create(ctx context.Context, action string) error {

	q := db.Query{
		QueryStr: "INSERT INTO logs (action) VALUES ($1)",
	}
	_, err := l.dbClient.DB().ContextExec(ctx, q, action)

	if err != nil {
		return err
	}
	return nil

}
