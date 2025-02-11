package chats

import (
	"chatservice/internal/repository"
	"chatservice/internal/service"
	"context"
)

type serv struct {
	repos repository.Repository
}

func NewService(repos repository.Repository) service.Service {
	return &serv{repos: repos}
}

func (s *serv) Create(ctx context.Context, usernames []string) (int, error) {
	return s.repos.Create(ctx, usernames)
}

func (s *serv) Delete(ctx context.Context, id int) error {
	return s.repos.Delete(ctx, id)
}
