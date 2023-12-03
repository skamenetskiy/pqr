package pqr

import (
	"context"
)

// BeforeUpdater interface.
type BeforeUpdater interface {
	// BeforeUpdate will be triggered right before "update" query is executed.
	// If an error is returned, the update will be canceled. On batch action,
	// the transaction will be aborted and rolled back.
	BeforeUpdate(context.Context) error
}

type BeforePutter interface {
	// BeforePut will be triggered right before "put" query is executed.
	// If an error is returned, the update will be canceled. On batch action,
	// the transaction will be aborted and rolled back.
	BeforePut(context.Context) error
}

type AfterFinder interface {
	AfterFind(ctx context.Context) error
}

func onBeforeUpdate[T any](ctx context.Context, t *T) error {
	if e, ok := any(t).(BeforeUpdater); ok {
		return e.BeforeUpdate(ctx)
	}
	return nil
}

func onBeforePut[T any](ctx context.Context, t *T) error {
	if e, ok := any(t).(BeforeUpdater); ok {
		return e.BeforeUpdate(ctx)
	}
	return nil
}

func onAfterFind[T any](ctx context.Context, t *T) error {
	if e, ok := any(t).(AfterFinder); ok {
		return e.AfterFind(ctx)
	}
	return nil
}
