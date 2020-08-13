package database

import (
	"kirby/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // use postgres dialect
	"github.com/lib/pq"
)

var (
	// Pg represents a database connection
	Pg *gorm.DB
)

const (
	// UniqueConstraintUserEmail enforces unique user emails
	UniqueConstraintUserEmail = "uix_users_email"
)

// PgConnect opens a connection to the database
func PgConnect() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", config.Env.PostgresURI)
	if err != nil {
		return nil, err
	}
	Pg = db
	return Pg, nil
}

// IsUniqueConstraintError determines whether an error originating from the database is
// the result of a unique constraint violation
func IsUniqueConstraintError(err error, constraintName string) bool {
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505" && pqErr.Constraint == constraintName
	}
	return false
}
