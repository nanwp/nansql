package nanlib

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// SQLServerConnectionManager is a struct that provides methods for managing connections to a SQL Server database.
type SQLServerConnectionManager struct {
	db *sqlx.DB
}

// NewConnectionManager creates a new instance of SQLServerConnectionManager and establishes a connection to the SQL Server database.
// It takes a DatabaseConfig as input and returns a pointer to SQLServerConnectionManager and an error.
func NewConnectionManager(cfg DatabaseConfig) (*SQLServerConnectionManager, error) {
	db, err := sqlx.Connect(cfg.Driver, cfg.DSN)
	if err != nil {
		log.Fatalf("failed to connect to sql server: %v", err)
		return nil, err
	}

	if cfg.MaxIdleConnections != 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConnections)
	}

	if cfg.MaxOpenConnections != 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConnections)
	}

	if cfg.MaxIdleDuration != 0 {
		db.SetConnMaxIdleTime(cfg.MaxIdleDuration)
	}

	if cfg.MaxLifeTimeDuration != 0 {
		db.SetConnMaxLifetime(cfg.MaxLifeTimeDuration)
	}

	log.Println("connected to sql server")

	return &SQLServerConnectionManager{
		db: db,
	}, nil
}

// Close closes the connection to the SQL Server database.
func (cm *SQLServerConnectionManager) Close() error {
	log.Println("closing sql server connection")
	return cm.db.Close()
}

// GetQuery returns a SingleInstruction instance for executing a single SQL query.
func (cm *SQLServerConnectionManager) GetQuery() *SingleInstruction {
	return NewSingleInstruction(cm.db)
}

// GetTransaction returns a MultiInstruction instance for executing multiple SQL queries within a transaction.
func (cm *SQLServerConnectionManager) GetTransaction() *MultiInstruction {
	return NewMultiInstruction(cm.db)
}
