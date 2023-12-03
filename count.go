package pqr

import (
	"context"
	"fmt"
)

type CountParams struct {
	Condition
}

func (r *repository[T, K]) Count(params ...CountParams) (int64, error) {
	return r.CountContext(context.Background(), params...)
}

func (r *repository[T, K]) CountContext(ctx context.Context, params ...CountParams) (int64, error) {
	return countContext(ctx, r.m, r.q, first[CountParams](params))
}

func (r *transaction[T, K]) Count(params ...CountParams) (int64, error) {
	return r.CountContext(context.Background(), params...)
}

func (r *transaction[T, K]) CountContext(ctx context.Context, params ...CountParams) (int64, error) {
	return countContext(ctx, r.m, r.q, first[CountParams](params))
}

func countContext(ctx context.Context, m *meta, q querier, params CountParams) (int64, error) {
	p := newPlaceholders()
	query := fmt.Sprintf(`SELECT COUNT("%s") FROM "%s"`, m.key.name, m.name)
	if params.Condition != nil && !params.Condition.isNil() {
		where, err := params.Condition.toSQL(p)
		if err != nil {
			return 0, err
		}
		query += fmt.Sprintf(` WHERE %s`, where)
	}
	count := int64(0)
	err := q.QueryRowContext(ctx, query, p.args()...).Scan(&count)
	return count, err
}
