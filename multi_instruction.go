package nansql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Transaction is an interface that defines the methods for managing transactions.
type Transaction interface {
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

// MultiInstruction is a struct that provides methods for executing multiple SQL instructions within a transaction.
type MultiInstruction struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

// NewMultiInstruction creates a new instance of MultiInstruction.
func NewMultiInstruction(db *sqlx.DB) *MultiInstruction {
	return &MultiInstruction{
		db: db,
		tx: nil,
	}
}

// Begin starts a new transaction.
func (t *MultiInstruction) Begin(ctx context.Context) error {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	t.tx = tx
	return nil
}

// Commit commits the transaction.
func (t *MultiInstruction) Commit(ctx context.Context) error {
	err := t.tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Rollback rolls back the transaction.
func (t *MultiInstruction) Rollback(ctx context.Context) error {
	err := t.tx.Rollback()
	if err != nil {
		return err
	}

	return nil
}

// Query executes a query that returns multiple rows.
func (t *MultiInstruction) Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return t.tx.QueryxContext(ctx, query, args...)
}

// QueryRow executes a query that returns a single row.
func (t *MultiInstruction) QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return t.tx.QueryRowxContext(ctx, query, args...)
}

// Exec executes a query that does not return any rows.
func (t *MultiInstruction) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return t.tx.ExecContext(ctx, query, args...)
}

// Prepare prepares a statement for execution.
func (t *MultiInstruction) Prepare(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return t.tx.PreparexContext(ctx, query)
}

// Select executes a query that selects multiple rows and stores the result in the provided slice.
func (t *MultiInstruction) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.tx.SelectContext(ctx, dest, query, args...)
}

// Get executes a query that selects a single row and stores the result in the provided struct.
func (t *MultiInstruction) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.tx.GetContext(ctx, dest, query, args...)
}

// CommitAndClose commits the transaction and closes it.
func (t *MultiInstruction) CommitAndClose(ctx context.Context) error {
	err := t.tx.Commit()
	if err != nil {
		return err
	}

	t.tx = nil
	return nil
}

// RollbackAndClose rolls back the transaction and closes it.
func (t *MultiInstruction) RollbackAndClose(ctx context.Context) error {
	err := t.tx.Rollback()
	if err != nil {
		return err
	}

	t.tx = nil
	return nil
}

// Rebind replaces placeholders in the query with the appropriate dialect-specific sequence.
func (t *MultiInstruction) Rebind(query string) string {
	return t.tx.Rebind(query)
}

// NamedExec executes a named query.
func (t *MultiInstruction) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return t.tx.NamedExecContext(ctx, query, arg)
}
