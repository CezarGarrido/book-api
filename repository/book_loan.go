package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/CezarGarrido/book-api/entity"
)

type bookLoanPostgres struct {
	db *sql.DB
}

func NewBookLoanPostgresRepo(db *sql.DB) *bookLoanPostgres {
	return &bookLoanPostgres{db}
}

func (bookLoanPg *bookLoanPostgres) Create(ctx context.Context, bookLoan *entity.BookLoan) (*entity.BookLoan, error) {
	query := `INSERT INTO public.book_loans
	          (book_id, from_user_id, to_user_id, lent_at, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6) 
			  RETURNING id`

	stmt, err := bookLoanPg.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	dateNow := time.Now()
	bookLoan.CreatedAt = dateNow
	bookLoan.UpdatedAt = dateNow

	err = stmt.QueryRowContext(ctx,
		bookLoan.BookID,
		bookLoan.FromUserID,
		bookLoan.ToUserID,
		bookLoan.LentAt,
		bookLoan.CreatedAt,
		bookLoan.UpdatedAt,
	).Scan(&bookLoan.ID)

	return bookLoan, err
}

func (bookLoanPg *bookLoanPostgres) ReturnBook(ctx context.Context, bookLoan *entity.BookLoan) (*entity.BookLoan, error) {
	query := "UPDATE public.book_loans SET returned_at=$1 WHERE from_user_id=$2 and id=$3;"
	stmt, err := bookLoanPg.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(ctx, bookLoan.ReturnedAt, bookLoan.FromUserID, bookLoan.ID)
	if err != nil {
		return nil, entity.ErrBookLoanFailedReturn
	}

	return bookLoan, nil
}

func (bookLoanPg *bookLoanPostgres) FindByFromUserAndBookID(ctx context.Context, fromUserID, bookID int64) (*entity.BookLoan, error) {
	query := `SELECT id, book_id, from_user_id, to_user_id, lent_at, returned_at, created_at, updated_at
	          FROM public.book_loans WHERE from_user_id=$1 and id=$2;`
	rows, err := bookLoanPg.fetch(ctx, query, fromUserID, bookID)
	if err != nil {
		return nil, err
	}
	if len(rows) > 0 {
		return rows[0], nil
	}

	return nil, entity.ErrBookLoanNotFoud
}

func (bookLoanPg *bookLoanPostgres) FindByToUserID(ctx context.Context, toUserID int64) ([]*entity.BookLoan, error) {
	query := `SELECT id, book_id, from_user_id, to_user_id, lent_at, returned_at, created_at, updated_at
	          FROM public.book_loans WHERE to_user_id=$1;`
	return bookLoanPg.fetch(ctx, query, toUserID)
}

func (bookLoanPg *bookLoanPostgres) FindByFromUserID(ctx context.Context, fromUserID int64) ([]*entity.BookLoan, error) {
	query := `SELECT id, book_id, from_user_id, to_user_id, lent_at, returned_at, created_at, updated_at
	          FROM public.book_loans WHERE from_user_id=$1;`
	return bookLoanPg.fetch(ctx, query, fromUserID)
}

func (bookLoanPg *bookLoanPostgres) fetch(ctx context.Context, query string, args ...interface{}) ([]*entity.BookLoan, error) {
	rows, err := bookLoanPg.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*entity.BookLoan, 0)
	for rows.Next() {
		bookLoan := new(entity.BookLoan)
		err := rows.Scan(
			&bookLoan.ID,
			&bookLoan.BookID,
			&bookLoan.FromUserID,
			&bookLoan.ToUserID,
			&bookLoan.LentAt,
			&bookLoan.ReturnedAt,
			&bookLoan.CreatedAt,
			&bookLoan.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, bookLoan)
	}
	return payload, nil
}
