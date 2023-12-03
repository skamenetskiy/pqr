package pqr

import (
	"context"
	"fmt"
	"strings"
)

type UpdateParams[T any] struct {
	Condition
	Columns []string
	Items   []*T
}

func (r *repository[T, K]) Update(params UpdateParams[T]) error {
	return r.UpdateContext(context.Background(), params)
}

func (r *repository[T, K]) UpdateContext(ctx context.Context, params UpdateParams[T]) error {
	if len(params.Items) == 1 {
		return updateColumnsContext[T, K](ctx, r.q, r.m, params.Condition, params.Columns, params.Items[0])
	}
	return r.TransactionContext(ctx, func(tx Transaction[T, K]) error {
		var err error
		for _, item := range params.Items {
			if err = updateColumnsContext[T, K](ctx, r.q, r.m, params.Condition, params.Columns, item); err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *transaction[T, K]) Update(params UpdateParams[T]) error {
	return r.UpdateContext(context.Background(), params)
}

func (r *transaction[T, K]) UpdateContext(ctx context.Context, params UpdateParams[T]) error {
	var err error
	for _, item := range params.Items {
		if err = updateColumnsContext[T, K](ctx, r.q, r.m, params.Condition, params.Columns, item); err != nil {
			return err
		}
	}
	return nil
}

func updateColumnsContext[T any, K Key](ctx context.Context, q querier, m *meta, c Condition, columns []string, item *T) error {
	if err := onBeforeUpdate[T](ctx, item); err != nil {
		return err
	}
	cm := make(map[string]struct{}, len(columns))
	if len(columns) > 0 {
		for _, cn := range columns {
			cm[cn] = struct{}{}
			if _, ok := m.columnsMap[cn]; !ok {
				return fmt.Errorf("invalid column %s", cn)
			}
		}
	} else {
		for _, cl := range m.columnsWithoutPK {
			cm[cl.name] = struct{}{}
			columns = append(columns, cl.name)
		}
	}
	p := newPlaceholders()
	pairs := make([]string, 0, len(m.columnsWithoutPK))
	values := newValueable[T](item, m.columnsWithPK).Values()
	if isZeroValue[K](values[m.key.index].(K)) {
		return fmt.Errorf("key cannot be unset on update")
	}
	for _, cl := range m.columnsWithoutPK {
		if _, ok := cm[cl.name]; !ok {
			continue
		}
		next := p.next(values[cl.index])
		pairs = append(pairs, fmt.Sprintf(`"%s" = %s`, cl.name, next[0]))
	}
	where, err := c.toSQL(p)
	if err != nil {
		return err
	}
	query := fmt.Sprintf(`UPDATE "%s" SET %s WHERE %s`,
		m.name, strings.Join(pairs, ", "), where)
	if err = q.QueryRowContext(ctx, query, p.args()...).Err(); err != nil {
		return err
	}
	return nil
}
