package chat

import (
	"chatservice/internal/repository"
	"context"

	db "github.com/quietdevil/Platform_common/pkg/db"
)

type repos struct {
	dbClient db.Client
}

func NewRepository(dbClient db.Client) repository.Repository {
	return &repos{dbClient: dbClient}
}

func (r *repos) Create(ctx context.Context, usernames []string) (int, error) {
	qs := "INSERT INTO chats (usernames) VALUES ($1) RETURNING id"

	q := db.Query{
		QueryStr: qs,
	}
	row, err := r.dbClient.DB().ContextQuery(ctx, q, usernames)
	if err != nil {
		return 0, err
	}
	var id int
	for row.Next() {
		row.Scan(&id)
	}

	return id, nil
}

func (r *repos) Delete(ctx context.Context, id int) error {
	qs := "DELETE FROM chats WHERE id=$1"

	q := db.Query{
		QueryStr: qs,
	}
	_, err := r.dbClient.DB().ContextExec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil

}
