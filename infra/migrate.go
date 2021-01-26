package infra

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func NewMigratePostgres(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	migrationDir := "./storage/migrations"

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationDir),
		"postgres", driver)
	if err != nil {
		return err
	}
	m.Steps(6)

	return err
}

func createPostgresDatabase() {
	db, err := sql.Open("postgres", newPgDSN())
	if err != nil {
		panic(err)
	}
	db.Exec("create database " + dbname)
	defer db.Close()
}

func newPgDSN() string {
	return fmt.Sprintf("user=%s password=%s host=%s sslmode=disable",
		user, password, host)
}
