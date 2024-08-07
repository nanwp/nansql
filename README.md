# nansql

`nansql` is a Golang library designed for managing connections to a SQL Server database using the `sqlx` package. It provides a simple and efficient way to handle database connections, execute queries, and manage transactions.

## Features

- Establishes and manages connections to a SQL Server database.
- Configurable connection pool settings.
- Supports query execution and transaction management.
- Provides an interface for different types of database operations.
- Allows seamless switching between regular queries and transactions using a unified interface.

## Installation

To install the `nansql` package, you can use `go get`:

```sh
go get github.com/nanwp/nansql
```

## Usage

Here's an example of how to use the `nansql` library with multiple repositories and transactions:

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/nanwp/nansql"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

func main() {
    cfg := nansql.DatabaseConfig{
        Driver:             "postgres", // you database driver dont forget to install driver postgre github.com/lib/pq
        DSN:                "your-dsn", // "postgres://nanda:nanda@localhost:5432/test?sslmode=disable"
        MaxIdleConnections: 10, 
        MaxOpenConnections: 100,
        MaxIdleDuration:    5 * time.Minute,
        MaxLifeTimeDuration: 1 * time.Hour,
    }

    manager, err := nansql.NewConnectionManager(cfg)
    if err != nil {
        log.Fatalf("Failed to create connection manager: %v", err)
    }
    defer manager.Close()

    ctx := context.Background()

    // Start a transaction
    tx := manager.GetTransaction()
    err = tx.Begin(ctx)
    if err != nil {
        log.Fatalf("Failed to begin transaction: %v", err)
    }

    // Initialize repositories with the transaction
    itemRepo := items.New(tx)

    // Perform operations within the transaction
    err = itemRepo.InsertItems(ctx, "Example Item")
    if err != nil {
        tx.Rollback(ctx)
        log.Fatalf("Failed to insert items: %v", err)
    }

    // Commit the transaction
    err = tx.Commit(ctx)
    if err != nil {
        log.Fatalf("Failed to commit transaction: %v", err)
    }

    log.Println("Transaction committed successfully")
}
```

```go
package items

import (
    "context"
    "github.com/nanwp/nansql"
)

type ItemsRepository struct {
    conn nansql.Connection
}

func New(conn nansql.Connection) *ItemsRepository {
    return &ItemsRepository{conn}
}

func (r *ItemsRepository) InsertItems(ctx context.Context, name string) error {
    query := `INSERT INTO items(name) VALUES ($1)`
    _, err := r.conn.Exec(ctx, query, name)
    if err != nil {
        return err
    }
    return nil
}
```

## Documentation

### SQLServerConnectionManager

`SQLServerConnectionManager` is a struct that provides methods for managing connections to a SQL Server database.

#### Methods

- `NewConnectionManager(cfg DatabaseConfig) (*SQLServerConnectionManager, error)`: Creates a new instance of `SQLServerConnectionManager` and establishes a connection to the SQL Server database.
- `Close() error`: Closes the connection to the SQL Server database.
- `GetQuery() *SingleInstruction`: Returns a `SingleInstruction` instance for executing a single SQL query.
- `GetTransaction() *MultiInstruction`: Returns a `MultiInstruction` instance for executing multiple SQL queries within a transaction.

### Connection Interface

`Connection` is an interface that represents a database connection.

#### Methods

- `Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)`: Executes a query that returns multiple rows.
- `QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row`: Executes a query that returns a single row.
- `Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)`: Executes a query that doesn't return any rows.
- `Prepare(ctx context.Context, query string) (*sqlx.Stmt, error)`: Prepares a query for execution.
- `Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error`: Executes a query that selects rows into a slice of structs or maps.
- `Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error`: Executes a query that selects a single row into a struct or map.
- `Rebind(query string) string`: Returns a query string with placeholders replaced with the appropriate dialect-specific sequence.
- `NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error)`: Executes a named query.

## Key Advantage

One of the key advantages of this library is the ability to seamlessly switch between using regular queries and transactions. This is facilitated by the unified interface, which both `GetQuery` and `GetTransaction` methods implement. This means that the repository only needs to work with the `Connection` interface, allowing for flexible and efficient database operations.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

---

Feel free to adjust any specific details to match your project requirements.