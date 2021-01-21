package infra

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	host     = "localhost"
	portDB   = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "blog_db"
)

type DB struct {
	SQL *sql.DB
}

func NewPostgres(dsn string) (*DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{SQL: db}, nil
}

func NewPostgresDSN() string {
	envUrl, ok := os.LookupEnv("DATABASE_URL")
	if ok {
		return envUrl
	}
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portDB, user, password, dbname)
}
