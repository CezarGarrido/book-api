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
	          (name, email, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4) 
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

func (userPg *userPostgres) FindUserByID(ctx context.Context, id int64) (*entity.User, error) {
	query := `SELECT id, "name", email, created_at, updated_at FROM public.users WHERE id=$1;`
	rows, err := userPg.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	if len(rows) > 0 {
		return rows[0], nil
	}

	return nil, entity.ErrUserNotFoud
}

func (userPg *userPostgres) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.User, error) {
	rows, err := userPg.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*entity.User, 0)
	for rows.Next() {
		user := new(entity.User)
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, user)
	}
	return payload, nil
}
