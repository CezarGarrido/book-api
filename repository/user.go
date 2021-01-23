package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/CezarGarrido/book-api/entity"
	"github.com/lib/pq"
)

type userPostgres struct {
	db *sql.DB
}

func NewUserPostgresRepo(db *sql.DB) *userPostgres {
	return &userPostgres{db}
}

func (userPg *userPostgres) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `INSERT INTO public.users 
	          (name, email, password, created_at, updated_at) 
			  VALUES ($1,$2,$3,$4,$5) 
			  RETURNING id`

	stmt, err := userPg.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	dateNow := time.Now()
	user.CreatedAt = dateNow
	user.UpdatedAt = dateNow

	err = stmt.QueryRowContext(ctx,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)

	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return nil, entity.ErrDuplicatedEmail
		}
		return nil, err
	}

	return user, nil
}

func (userPg *userPostgres) FindAllUsers(ctx context.Context) ([]*entity.User, error) {
	return nil, nil
}
