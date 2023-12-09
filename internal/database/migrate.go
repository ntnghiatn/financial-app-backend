package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattes/migrate/source/file"
	"github.com/ntnghiatn/financial-app-backend/internal/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func migrateDb(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "connecting to database")
	}

	migrationSource := fmt.Sprintf("file://%sinternal/database/migrations/", *config.DataDirectory)
	//logrus.Info(migrationSource)
	migrator, err := migrate.NewWithDatabaseInstance(migrationSource, "ciquandb", driver)
	logrus.Info(migrator)
	if err != nil {
		return errors.Wrap(err, "creating migrator")
	}

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "executing migration")
	}

	version, dirty, err := migrator.Version()
	if err != nil {
		return errors.Wrap(err, "getting migration version")
	}

	logrus.WithFields(logrus.Fields{"version": version, "dirty": dirty}).Debug("Database migrated!!")

	return nil
}
