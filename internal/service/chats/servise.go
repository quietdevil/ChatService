package chats

import (
	"context"
	"fmt"
	"github.com/quietdevil/ChatSevice/internal/repository"
	"github.com/quietdevil/ChatSevice/internal/service"

	db "github.com/quietdevil/Platform_common/pkg/db"
)

type serv struct {
	repos repository.Repository
	tx    db.TxManager
	logs  repository.Logger
}

func NewService(repos repository.Repository, tx db.TxManager, logs repository.Logger) service.Service {
	return &serv{repos: repos,
		tx: tx, logs: logs}
}

func (s *serv) Create(ctx context.Context, usernames []string) (int, error) {
	var id int
	err := s.tx.ReadCommited(ctx, func(ctx context.Context) error {
		idD, err := s.repos.Create(ctx, usernames)
		if err != nil {
			return err
		}
		action := fmt.Sprintf("Create chats with users: %v", usernames)
		err = s.logs.Create(ctx, action)
		if err != nil {
			return err
		}
		id = idD

		return nil

	})

	return id, err
}

func (s *serv) Delete(ctx context.Context, id int) error {
	return s.repos.Delete(ctx, id)
}
