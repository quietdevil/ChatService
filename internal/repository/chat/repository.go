package chat

import (
	"chatservice/internal/repository"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repos struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.Repository {
	return &repos{db: db}
}

func (r *repos) Create(ctx context.Context, usernames []string) (int, error) {
	q := "INSERT INTO chats (usernames) VALUES ($1) RETURNING id"
	row, err := r.db.Query(ctx, q, usernames)
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
	q := "DELETE FROM chats WHERE id=$1"
	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil

}
