package pqr

import (
	"context"
	"database/sql"
	"time"
)

type SQL interface {
	querier
	txBeginner
}

type txBeginner interface {
	BeginTx(ctx context.Context) (Tx, error)
}

type querier interface {
	QueryContext(ctx context.Context, query string, args ...any) (Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) Row
}

type Tx interface {
	querier
	Commit() error
	Rollback() error
}

type Rows interface {
	Row
	Next() bool
	Close() error
}

type Row interface {
	Err() error
	Scan(...any) error
}

type sqlDbWrapper struct {
	db *sql.DB
}

func (s *sqlDbWrapper) BeginTx(ctx context.Context) (Tx, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	return &sqlDbTxWrapper{tx}, err
}

func (s *sqlDbWrapper) QueryContext(ctx context.Context, query string, args ...any) (Rows, error) {
	return s.db.QueryContext(ctx, query, args...)
}

func (s *sqlDbWrapper) QueryRowContext(ctx context.Context, query string, args ...any) Row {
	return s.db.QueryRowContext(ctx, query, args...)
}

type sqlDbTxWrapper struct {
	tx *sql.Tx
}

func (s *sqlDbTxWrapper) QueryContext(ctx context.Context, query string, args ...any) (Rows, error) {
	return s.tx.QueryContext(ctx, query, args...)
}

func (s *sqlDbTxWrapper) QueryRowContext(ctx context.Context, query string, args ...any) Row {
	return s.tx.QueryRowContext(ctx, query, args...)
}

func (s *sqlDbTxWrapper) Commit() error {
	return s.tx.Commit()
}

func (s *sqlDbTxWrapper) Rollback() error {
	return s.tx.Rollback()
}

// queryLogger is a wrapper around SQL that logs queries and its duration.
type queryLogger struct {
	q SQL
	l Logger
}

// BeginTx just proxies the call to SQL.
func (e *queryLogger) BeginTx(ctx context.Context) (Tx, error) {
	return e.q.BeginTx(ctx)
}

// QueryContext logs the query before and after the query is executed.
func (e *queryLogger) QueryContext(ctx context.Context, query string, args ...any) (Rows, error) {
	e.l.BeforeQuery(ctx, query, args)
	ts := time.Now()
	rows, err := e.q.QueryContext(ctx, query, args...)
	e.l.AfterQuery(ctx, query, args, err, time.Since(ts))
	return rows, err
}

// QueryRowContext logs the query before and after the query is executed.
func (e *queryLogger) QueryRowContext(ctx context.Context, query string, args ...any) Row {
	e.l.BeforeQuery(ctx, query, args)
	ts := time.Now()
	row := e.q.QueryRowContext(ctx, query, args...)
	e.l.AfterQuery(ctx, query, args, row.Err(), time.Since(ts))
	return row
}
