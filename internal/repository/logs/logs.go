package logs

import (
	"context"
	"github.com/quietdevil/ChatSevice/internal/repository"

	db "github.com/quietdevil/Platform_common/pkg/db"
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
