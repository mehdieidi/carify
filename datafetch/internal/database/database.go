//go:build !postgres
// +build !postgres

package database

import "database/sql"

func Open(dsn string) (*sql.DB, error) {
	panic("use a database")
}

func AutoMigrate(db *sql.DB) error {
	panic("use a database")
}
