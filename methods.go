package pqr

import (
	"database/sql"
	"errors"
)

// New is a Repository factory, where name is the database table name and queryLogger
// is the database connection that implements querier interface.
func New[T any, K Key](name string, opts ...Option) Repository[T, K] {
	c := &config{
		l: NopLogger,
	}
	for _, opt := range opts {
		opt.apply(c)
	}

	if c.q == nil {
		panic("cannot init without a db connection")
	}

	m, err := parseMeta(name, new(T))
	if err != nil {
		panic(err)
	}

	r := &repository[T, K]{
		q: &queryLogger{c.q, c.l},
		m: m,
		l: c.l,
	}
	if r.m.key.name == "" {
		panic("cannot init repository without a key")
	}
	return r
}

// IsNotFound is a helper functions to check if the error means that no records were found.
func IsNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
