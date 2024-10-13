package database

import (
	"context"
	"fmt"
	"github.com/nopecho/golang-template/internal/app/domain"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type GormTransaction struct {
	tx *gorm.DB
}

func NewGormTransaction(tx *gorm.DB) *GormTransaction {
	return &GormTransaction{tx: tx}
}

func (t *GormTransaction) GetTx() any {
	return t.tx
}

type GormTransactionManager struct {
	db *gorm.DB
}

const (
	txTimeout = 2 * time.Second
)

func NewGormTransactionManager(db *gorm.DB) *GormTransactionManager {
	return &GormTransactionManager{db: db}
}

func (m *GormTransactionManager) Execute(ctx context.Context, stat domain.TxFunc) error {
	return m.execute(ctx, txTimeout, stat)
}

func (m *GormTransactionManager) ExecuteWithTimeout(ctx context.Context, timeout time.Duration, stat domain.TxFunc) error {
	return m.execute(ctx, timeout, stat)
}

func (m *GormTransactionManager) execute(ctx context.Context, timeout time.Duration, stat domain.TxFunc) error {
	cancelCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	err := m.db.Transaction(func(tx *gorm.DB) error {
		select {
		case <-cancelCtx.Done():
			return errors.Cause(fmt.Errorf("transaction timeout after %fs", timeout.Seconds()))
		default:
			err := stat(NewGormTransaction(tx))
			if err != nil {
				return err
			}
			return nil
		}
	})
	return err
}
