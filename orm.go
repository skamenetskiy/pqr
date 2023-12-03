package pqr

import (
	"context"
)

type orm[T any, K Key] interface {
	// Count is a shortcut to CountContext with context.Background().
	Count(params ...CountParams) (int64, error)

	// CountContext counts the number of elements in the repository.
	CountContext(ctx context.Context, params ...CountParams) (int64, error)

	// FindOne is a shortcut to FindOneContext with context.Background().
	FindOne(id K) (*T, error)

	// FindOneContext finds one element in the repository by key.
	FindOneContext(ctx context.Context, key K) (*T, error)

	// Find is a shortcut to FindContext with context.Background().
	Find(params ...FindParams) ([]*T, error)

	// FindContext searches for items in Repository by FindParams.
	FindContext(ctx context.Context, params ...FindParams) ([]*T, error)

	Put(items ...*T) error
	PutContext(ctx context.Context, items ...*T) error

	Delete(params ...DeleteParams) error
	DeleteContext(ctx context.Context, params ...DeleteParams) error

	Update(params UpdateParams[T]) error
	UpdateContext(ctx context.Context, params UpdateParams[T]) error
}
