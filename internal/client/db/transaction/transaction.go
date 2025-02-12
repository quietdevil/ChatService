package transaction

import (
	"chatservice/internal/client/db"
	"chatservice/internal/client/db/pg"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type TxManager struct {
	db db.Transactor
}

func NewManager(db db.Transactor) db.TxManager {
	return &TxManager{
		db: db,
	}
}

func (m *TxManager) transaction(ctx context.Context, opt pgx.TxOptions, handler db.Hadler) (err error) {

	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return handler(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opt)
	if err != nil {
		return errors.Wrap(err, "can't begin transaction")
	}

	ctx = pg.MakeContext(ctx, tx)

	defer func() {
		// восстанавливаемся после паники
		if r := recover(); r != nil {
			err = errors.Errorf("panic recovered: %v", r)
		}

		// откатываем транзакцию, если произошла ошибка
		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "errRollback: %v", errRollback)
			}

			return
		}

		// если ошибок не было, коммитим транзакцию
		if nil == err {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "tx commit failed")
			}
		}
	}()

	// Выполните код внутри транзакции.
	// Если функция терпит неудачу, возвращаем ошибку, и функция отсрочки выполняет откат
	// или в противном случае транзакция коммитится.
	if err = handler(ctx); err != nil {
		err = errors.Wrap(err, "failed executing code inside transaction")
	}
	return err
}

func (m *TxManager) ReadCommited(ctx context.Context, handler db.Hadler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.transaction(ctx, txOpts, handler)
}
