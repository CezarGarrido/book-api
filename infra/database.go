package infra

import (
	"database/sql"
	"fmt"
)

const (
	host       = "localhost"
	portDB     = "5432"
	user       = "postgres"
	password   = "postgres"
	dbname     = "books_db"
	dbTestName = "books_db_test"
)

type DB struct {
	SQL *sql.DB
}

func NewPostgres(dsn string) (*DB, error) {
	createPostgresDatabase()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &DB{SQL: db}, nil
}

func NewPostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portDB, user, password, dbname)
}

func NewPostgresTestDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portDB, user, password, dbTestName)
}
