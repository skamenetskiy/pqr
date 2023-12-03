package pqr

import (
	"context"
	"fmt"
	"reflect"
	"strings"
)

func (r *repository[T, K]) Put(items ...*T) error {
	return r.PutContext(context.Background(), items...)
}

func (r *repository[T, K]) PutContext(ctx context.Context, items ...*T) error {
	if len(items) == 0 {
		return nil
	}
	return r.Transaction(func(tx Transaction[T, K]) error {
		for _, item := range items {
			if err := putContext[T, K](ctx, r.m, tx, item); err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *transaction[T, K]) Put(items ...*T) error {
	return r.PutContext(context.Background(), items...)
}

func (r *transaction[T, K]) PutContext(ctx context.Context, items ...*T) error {
	for _, item := range items {
		if err := putContext[T, K](ctx, r.m, r.q, item); err != nil {
			return err
		}
	}
	return nil
}

func putContext[T any, K Key](ctx context.Context, m *meta, q querier, t *T) error {
	if err := onBeforePut[T](ctx, t); err != nil {
		return err
	}
	values := newValueable[T](t, m.columnsWithPK).Values()
	columns := m.columnsWithPK
	if isZeroValue[K](values[m.key.index].(K)) {
		values = withoutIndex(values, m.key.index)
		columns = m.columnsWithoutPK
	}
	p := newPlaceholders()
	query := fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s) RETURNING "%s"`,
		m.name, columnsToString(columns), strings.Join(p.next(values...), ", "), m.key.name)
	pk := reflect.ValueOf(t).Elem().Field(m.key.index).Addr().Interface()
	if err := q.QueryRowContext(ctx, query, p.args()...).Scan(pk); err != nil {
		return err
	}
	return nil
}
