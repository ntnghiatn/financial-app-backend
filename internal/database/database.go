package database

import (
	"io"

	"github.com/jmoiron/sqlx"
)

// UniqueViolation Postgres Error string for a unique index violation
const UniqueViolation = "unique_violation"

// Database - interface
type Database interface {
	UsersDB

	io.Closer
}

type database struct {
	conn *sqlx.DB
}

func (d *database) Close() error {
	return d.conn.Close()
}
