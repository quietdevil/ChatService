package service

import "context"

type Service interface {
	Create(ctx context.Context, usernames []string) (int, error)
	Delete(ctx context.Context, id int) error
}
