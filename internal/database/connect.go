package database

import (
	"flag"
	"time"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var (
	databaseURL     = flag.String("database-url", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable", "Database URL")
	databaseTimeout = flag.Int64("database-timeout-ms", 2000, "")
)

func Connect() (*sqlx.DB, error) {
	//Connect to database
	dbURL := *databaseURL

	logrus.WithField("url", dbURL).Debug("Connecting to database...")
	conn, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to database")
	}

	conn.SetMaxIdleConns(32)

	//Check if database running
	if err := waitForDB(conn); err != nil {
		return nil, err
	}

	//migrate database schema
	if err := migrateDb(conn.DB); err != nil {
		return nil, errors.Wrap(err, "could not migrate database")
	}

	return conn, nil
}

// New - creates new database
func New() (Database, error) {
	conn, err := Connect()
	if err != nil {
		return nil, err
	}
	d := &database{conn: conn}
	return d, nil
}

func waitForDB(conn *sqlx.DB) error {
	ready := make(chan struct{})
	go func() {
		for {
			if err := conn.Ping(); err != nil {
				close(ready)
				return
			}
			time.Sleep(100 * time.Microsecond)
		}
	}()

	select {
	case <-ready:
		return nil
	case <-time.After(time.Duration((*databaseTimeout) * int64(time.Millisecond))):
		return errors.New("Database not ready")
	}
}
