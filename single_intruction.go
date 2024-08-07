// Package nanlib provides a set of utilities for working with databases using the nanlib library.
package nanlib

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// SingleInstruction represents a set of methods for executing single database instructions.
type SingleInstruction struct {
	db *sqlx.DB
}

// NewSingleInstruction creates a new instance of SingleInstruction.
func NewSingleInstruction(db *sqlx.DB) *SingleInstruction {
	return &SingleInstruction{
		db: db,
	}
}

// Query executes a query that returns rows.
// It takes a context, a query string, and optional arguments.
// It returns a *sqlx.Rows and an error.
func (s *SingleInstruction) Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return s.db.QueryxContext(ctx, query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
// It takes a context, a query string, and optional arguments.
// It returns a *sqlx.Row.
func (s *SingleInstruction) QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return s.db.QueryRowxContext(ctx, query, args...)
}

// Exec executes a query that doesn't return any rows.
// It takes a context, a query string, and optional arguments.
// It returns a sql.Result and an error.
func (s *SingleInstruction) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return s.db.ExecContext(ctx, query, args...)
}

// Prepare prepares a statement for later queries or executions.
// It takes a context and a query string.
// It returns a *sqlx.Stmt and an error.
func (s *SingleInstruction) Prepare(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return s.db.PreparexContext(ctx, query)
}

// Select executes a query that returns multiple rows and stores the result in the provided slice.
// It takes a context, a destination slice, a query string, and optional arguments.
// It returns an error.
func (s *SingleInstruction) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return s.db.SelectContext(ctx, dest, query, args...)
}

// Get executes a query that is expected to return at most one row and stores the result in the provided struct.
// It takes a context, a destination struct, a query string, and optional arguments.
// It returns an error.
func (s *SingleInstruction) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return s.db.GetContext(ctx, dest, query, args...)
}

// Rebind takes a query string and returns a new query string with all occurrences of '?' replaced with the database's bindvar syntax.
func (s *SingleInstruction) Rebind(query string) string {
	return s.db.Rebind(query)
}

// NamedExec executes a named query using the provided struct or map as the named arguments.
// It takes a context, a query string, and an argument.
// It returns a sql.Result and an error.
func (s *SingleInstruction) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return s.db.NamedExecContext(ctx, query, arg)
}
