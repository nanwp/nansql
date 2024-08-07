package nanlib

import "time"

// DatabaseConfig is a struct that contains the configuration for connecting to a database.
type DatabaseConfig struct {
	Driver              string
	DSN                 string
	MaxIdleConnections  int
	MaxOpenConnections  int
	MaxIdleDuration     time.Duration
	MaxLifeTimeDuration time.Duration
}
