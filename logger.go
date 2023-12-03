package pqr

import (
	"context"
	"log/slog"
	"time"
)

// Logger interface for query logging.
type Logger interface {
	// BeforeQuery is triggered before the query is executed.
	BeforeQuery(ctx context.Context, query string, args []any)

	// AfterQuery is triggered right after the query was executed.
	AfterQuery(ctx context.Context, query string, args []any, err error, d time.Duration)
}

const (
	// NopLogger is a Logger that does nothing. It is used by default,
	// if no logger is provided on initialization.
	NopLogger = nopLogger('🚀')

	// DefaultLogger is a wrapper around slog.Logger.
	DefaultLogger = slogLogger('🚀')
)

type nopLogger rune

func (nopLogger) BeforeQuery(_ context.Context, _ string, _ []any) {}

func (nopLogger) AfterQuery(_ context.Context, _ string, _ []any, _ error, _ time.Duration) {}

type slogLogger rune

func (slogLogger) BeforeQuery(_ context.Context, query string, args []any) {
	slog.Debug("BeforeQuery", "query", query, "args", args)
}

func (slogLogger) AfterQuery(_ context.Context, query string, args []any, err error, d time.Duration) {
	if err != nil {
		slog.Error("", "query", query, "args", args, "duration", d.String(), "error", err)
		return
	}
	slog.Debug("", "query", query, "args", args, "duration", d.String())
}
