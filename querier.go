package pqr

import (
	"context"
)

func (r *repository[T, K]) QueryContext(ctx context.Context, query string, args ...any) (Rows, error) {
	return r.q.QueryContext(ctx, query, args...)
}

func (r *repository[T, K]) QueryRowContext(ctx context.Context, query string, args ...any) Row {
	return r.q.QueryRowContext(ctx, query, args...)
}
