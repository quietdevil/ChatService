package chats

import (
	"chatservice/internal/client/db"
	"chatservice/internal/repository"
	"chatservice/internal/service"
	"context"
)

type serv struct {
	repos repository.Repository
	tx    db.TxManager
}

func NewService(repos repository.Repository, tx db.TxManager) service.Service {
	return &serv{repos: repos,
		tx: tx}
}

func (s *serv) Create(ctx context.Context, usernames []string) (int, error) {
	return s.repos.Create(ctx, usernames)
}

func (s *serv) Delete(ctx context.Context, id int) error {
	return s.repos.Delete(ctx, id)
}
