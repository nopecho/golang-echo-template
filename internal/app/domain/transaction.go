package domain

import (
	"context"
	"time"
)

type TxFunc func(tx Transaction) error

type TransactionManager interface {
	// Execute executes a transaction
	Execute(ctx context.Context, stat TxFunc) error
	// ExecuteWithTimeout executes a transaction with a timeout
	ExecuteWithTimeout(ctx context.Context, timeout time.Duration, stat TxFunc) error
}

type Transaction interface {
	GetTx() any
}
