package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/CezarGarrido/book-api/entity"
)

type bookPostgres struct {
	db *sql.DB
}

func NewBookPostgresRepo(db *sql.DB) *bookPostgres {
	return &bookPostgres{db}
}

func (bookPg *bookPostgres) Create(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	query := `INSERT INTO public.books 
	          (user_id, title, pages, author_name, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6) 
			  RETURNING id`

	stmt, err := bookPg.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	dateNow := time.Now()
	book.CreatedAt = dateNow
	book.UpdatedAt = dateNow

	err = stmt.QueryRowContext(ctx,
		book.UserID,
		book.Title,
		book.Pages,
		book.AuthorName,
		book.CreatedAt,
		book.UpdatedAt,
	).Scan(&book.ID)

	return book, err
}

func (bookPg *bookPostgres) FindByID(ctx context.Context, bookID int64) (*entity.Book, error) {
	query := `SELECT id, user_id, title, pages, author_name, created_at, updated_at FROM public.books WHERE id=$1;`
	rows, err := bookPg.fetch(ctx, query, bookID)
	if err != nil {
		return nil, err
	}
	if len(rows) > 0 {
		return rows[0], nil
	}

	return nil, entity.ErrBookNotFoud
}

func (bookPg *bookPostgres) FindByUserID(ctx context.Context, userID int64) ([]*entity.Book, error) {
	query := `SELECT id, user_id, title, pages, author_name, created_at, updated_at FROM public.books WHERE id=$1;`
	return bookPg.fetch(ctx, query, userID)
}

func (bookPg *bookPostgres) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.Book, error) {
	rows, err := bookPg.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*entity.Book, 0)
	for rows.Next() {
		book := new(entity.Book)
		err := rows.Scan(
			&book.ID,
			&book.UserID,
			&book.Title,
			&book.Pages,
			&book.AuthorName,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, book)
	}
	return payload, nil
}
