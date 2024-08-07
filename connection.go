package nansql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Connection is an interface that represents a database connection.
type Connection interface {
	// Query executes a query that returns multiple rows.
	// It takes a context, the query string, and optional arguments.
	// It returns a *sqlx.Rows and an error.
	Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)

	// QueryRow executes a query that returns a single row.
	// It takes a context, the query string, and optional arguments.
	// It returns a *sqlx.Row.
	QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row

	// Exec executes a query that doesn't return any rows.
	// It takes a context, the query string, and optional arguments.
	// It returns a sql.Result and an error.
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// Prepare prepares a query for execution.
	// It takes a context and the query string.
	// It returns a *sqlx.Stmt and an error.
	Prepare(ctx context.Context, query string) (*sqlx.Stmt, error)

	// Select executes a query that selects rows into a slice of structs or maps.
	// It takes a context, a destination slice, the query string, and optional arguments.
	// It returns an error.
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// Get executes a query that selects a single row into a struct or map.
	// It takes a context, a destination struct or map, the query string, and optional arguments.
	// It returns an error.
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// Rebind returns a query string with placeholders replaced with the appropriate dialect-specific sequence.
	// It takes the query string and returns the modified query string.
	Rebind(query string) string

	// NamedExec executes a named query.
	// It takes a context, the query string, and a struct or map containing named arguments.
	// It returns a sql.Result and an error.
	NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}
