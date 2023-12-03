package pqr

import (
	"context"
	"fmt"
)

func (r *repository[T, K]) FindOne(id K) (*T, error) {
	return r.FindOneContext(context.Background(), id)
}

func (r *repository[T, K]) FindOneContext(ctx context.Context, key K) (*T, error) {
	return findOneContext[T, K](ctx, r.m, r.q, key)
}

func (r *transaction[T, K]) FindOne(id K) (*T, error) {
	return r.FindOneContext(context.Background(), id)
}

func (r *transaction[T, K]) FindOneContext(ctx context.Context, key K) (*T, error) {
	return findOneContext[T, K](ctx, r.m, r.q, key)
}

func findOneContext[T any, K Key](ctx context.Context, m *meta, q querier, key K) (*T, error) {
	query := fmt.Sprintf(`SELECT %s FROM "%s" WHERE "%s" = $1`,
		columnsToString(m.columnsWithPK), m.name, m.key.name)
	row := q.QueryRowContext(ctx, query, key)
	if err := row.Err(); err != nil {
		return nil, err
	}
	t := new(T)
	if err := row.Scan(newPointable[T](t, m.columnsWithPK).Pointers()...); err != nil {
		return nil, err
	}
	if err := onAfterFind[T](ctx, t); err != nil {
		return nil, err
	}
	return t, nil
}
