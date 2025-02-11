package repository

import "context"

type Repository interface {
	Create(ctx context.Context, usernames []string) (int, error)
	Delete(ctx context.Context, id int) error
}
