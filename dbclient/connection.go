package dbclient

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // use postgres dialect
	"github.com/lib/pq"
)

// Connection wraps the underlying database connection
type Connection struct {
	*gorm.DB
}

const (
	// UniqueConstraintUserEmail enforces unique user emails
	UniqueConstraintUserEmail = "uix_users_email"
)

// Connect opens a connection to the database
func Connect(dataSourceName string) (*Connection, error) {
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &Connection{db}, nil
}

// IsUniqueConstraintError determines whether an error originating from the database is
// the result of a unique constraint violation
func IsUniqueConstraintError(err error, constraintName string) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505" && pqErr.Constraint == constraintName
	}
	return false
}
