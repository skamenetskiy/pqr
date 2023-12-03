package pqr

import (
	"context"
	"fmt"
)

type FindParams struct {
	Condition
	Order
	OrderBy string
	Limit   int64
}

func (r *repository[T, K]) Find(params ...FindParams) ([]*T, error) {
	return r.FindContext(context.Background(), params...)
}

func (r *repository[T, K]) FindContext(ctx context.Context, params ...FindParams) ([]*T, error) {
	return findContext[T](ctx, r.m, r.q, first[FindParams](params))
}

func (r *transaction[T, K]) Find(params ...FindParams) ([]*T, error) {
	return r.FindContext(context.Background(), params...)
}

func (r *transaction[T, K]) FindContext(ctx context.Context, params ...FindParams) ([]*T, error) {
	return findContext[T](ctx, r.m, r.q, first[FindParams](params))
}

func findContext[T any](ctx context.Context, m *meta, q querier, params FindParams) ([]*T, error) {
	p := newPlaceholders()
	query := fmt.Sprintf(`SELECT %s FROM "%s"`, columnsToString(m.columnsWithPK), m.name)
	if params.Condition != nil && !params.Condition.isNil() {
		cs, err := params.Condition.toSQL(p)
		if err != nil {
			return nil, err
		}
		query += fmt.Sprintf(` WHERE %s`, cs)
	}
	if params.Order != OrderNone {
		var orderBy string
		if params.OrderBy != "" {
			orderBy = params.OrderBy
		} else {
			orderBy = m.key.name
		}
		query += fmt.Sprintf(` ORDER BY "%s" %s`, orderBy, params.Order)
	}
	if params.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %s", p.nextAny(params.Limit)...)
	}
	rows, err := q.QueryContext(ctx, query, p.args()...)
	if err != nil {
		if IsNotFound(err) {
			return make([]*T, 0, 1), nil
		}
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	res := make([]*T, 0)
	for rows.Next() {
		t := new(T)
		err = rows.Scan(newPointable[T](t, m.columnsWithPK).Pointers()...)
		if err != nil {
			return nil, err
		}
		if err = onAfterFind[T](ctx, t); err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}
