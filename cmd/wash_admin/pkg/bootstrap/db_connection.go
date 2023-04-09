package bootstrap

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gocraft/dbr/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"time"
)

func NewDbConn(db DBConfig) (dbPool *dbr.Connection, err error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		db.User, db.Password, db.Database, db.Host, db.Port)

	t := time.After(time.Second * 10)

	dbPool, err = dbr.Open("postgres", dsn, nil)
	if err != nil {
		return nil, err
	}

loop:
	for {
		select {
		case <-t:
			return
		default:
			err = dbPool.Ping()
			if err != nil {
				time.Sleep(time.Millisecond * 100)

				continue
			}
			break loop
		}
	}

	// Maximum amount of concurrent db connections
	dbPool.SetMaxOpenConns(10)
	dbPool.SetMaxOpenConns(5)
	dbPool.SetConnMaxLifetime(time.Second * 20)
	dbPool.SetConnMaxIdleTime(time.Second * 60)

	return
}

func UpMigrations(db *sql.DB, dbName string, migrationsPath string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		dbName, driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}

	return nil
}
