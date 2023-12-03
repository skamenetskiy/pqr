package pqr

import (
	"context"
)

type Transaction[T any, K Key] interface {
	orm[T, K]
	querier
}

type transaction[T any, K Key] struct {
	q querier
	m *meta
}

func (r *transaction[T, K]) QueryContext(ctx context.Context, query string, args ...any) (Rows, error) {
	return r.q.QueryContext(ctx, query, args...)
}

func (r *transaction[T, K]) QueryRowContext(ctx context.Context, query string, args ...any) Row {
	return r.q.QueryRowContext(ctx, query, args...)
}

type transactioner[T any, K Key] interface {
	Transaction(fn func(tx Transaction[T, K]) error) error
	TransactionContext(ctx context.Context, fn func(tx Transaction[T, K]) error) error
}

func (r *repository[T, K]) Transaction(fn func(tx Transaction[T, K]) error) error {
	return r.TransactionContext(context.Background(), fn)
}

func (r *repository[T, K]) TransactionContext(ctx context.Context, fn func(tx Transaction[T, K]) error) error {
	tx, err := r.q.BeginTx(ctx)
	if err != nil {
		return err
	}
	if err = fn(&transaction[T, K]{tx, r.m}); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return err2
		}
		return err
	}
	return tx.Commit()
}
