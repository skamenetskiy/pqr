package pqr

import (
	"context"
	"fmt"
)

type DeleteParams struct {
	Condition
}

func (r *repository[T, K]) Delete(params ...DeleteParams) error {
	return r.DeleteContext(context.Background(), params...)
}

func (r *repository[T, K]) DeleteContext(ctx context.Context, params ...DeleteParams) error {
	return deleteContext(ctx, r.m, r.q, first[DeleteParams](params))
}

func (r *transaction[T, K]) Delete(params ...DeleteParams) error {
	return r.DeleteContext(context.Background(), params...)
}

func (r *transaction[T, K]) DeleteContext(ctx context.Context, params ...DeleteParams) error {
	return deleteContext(ctx, r.m, r.q, first[DeleteParams](params))
}

func deleteContext(ctx context.Context, m *meta, q querier, params DeleteParams) error {
	p := newPlaceholders()
	query := fmt.Sprintf(`DELETE FROM "%s"`, m.name)
	if params.Condition != nil && !params.Condition.isNil() {
		where, err := params.Condition.toSQL(p)
		if err != nil {
			return err
		}
		query += fmt.Sprintf(` WHERE %s`, where)
	}
	if err := q.QueryRowContext(ctx, query, p.args()...).Err(); err != nil {
		return err
	}
	return nil
}
