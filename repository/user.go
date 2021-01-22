package repository

import (
	"context"
	"database/sql"

	"github.com/CezarGarrido/book-api/entity"
)

type userPostgres struct {
	db *sql.DB
}

func NewUserPostgresRepo(db *sql.DB) *userPostgres {
	return &userPostgres{db}
}

func (userPg *userPostgres) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	return nil, nil
}
