package pqr

import (
	"database/sql"
)

// Option interface is used for initial configuration.
type Option interface{ apply(optional) }

// optional interface describes a set of methods for initialization.
type optional interface {
	setLogger(l Logger)
	setQuerier(q SQL)
}

// WithLogger attaches a Logger to Repository on initialization.
func WithLogger(l Logger) Option {
	return &commonOption[Logger]{func(o optional) {
		o.setLogger(l)
	}}
}

// WithStandardSQL attaches sql.DB connection to Repository on initialization.
func WithStandardSQL(db *sql.DB) Option {
	return &commonOption[*sql.DB]{func(o optional) {
		o.setQuerier(&sqlDbWrapper{db})
	}}
}

// WithSQL attaches SQL connection to Repository on initialization.
func WithSQL(s SQL) Option {
	return &commonOption[SQL]{func(o optional) {
		o.setQuerier(s)
	}}
}

type commonOption[T any] struct {
	f func(optional)
}

func (o *commonOption[T]) apply(r optional) { o.f(r) }
