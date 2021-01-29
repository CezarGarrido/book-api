package entity

import (
	"context"
	"errors"
	"time"
)

var ErrBookLoanNotFoud = errors.New("Não foi possível recuperar o empréstimo")
var ErrBookLoanIsReturned = errors.New("Livro já devolvido")
var ErrBookLoanFailedReturn = errors.New("Não foi possível devolver o livro")

// BookLoan represents a loan for a book
type BookLoan struct {
	ID         int64      `json:"id"`          // Loan id
	BookID     int64      `json:"book_id"`     // Book id
	FromUserID int64      `json:"from_user"`   // FromUserID represents the id of the user who is lending a book
	ToUserID   int64      `json:"to_user"`     // Represents the id of the user who is going to borrow the book
	LentAt     time.Time  `json:"lent_at"`     // Represents the date the loan was made
	ReturnedAt *time.Time `json:"returned_at"` // Represents the date the book was returned - If the value is equal to null, it means that the book has not yet been returned
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type BookLoanUsecase interface {
	// Lend Book
	LendBook(ctx context.Context, user User, toUserID int64, bookID int64) (*BookLoan, error)
	// Return book
	ReturnBook(ctx context.Context, user User, bookID int64) (*BookLoan, error)
	// return the books I lent to someone
	FindByFromUserID(ctx context.Context, fromUserID int64) ([]*BookLoan, error)
	// Return books I borrowed from someone
	FindByToUserID(ctx context.Context, toUserID int64) ([]*BookLoan, error)
}

type BookLoanRepo interface {
	// Create loan in the database
	Create(ctx context.Context, bookLoan *BookLoan) (*BookLoan, error)
	// Return book
	ReturnBook(ctx context.Context, bookLoan *BookLoan) (*BookLoan, error)
	// Find by book ID
	FindByFromUserAndBookID(ctx context.Context, fromUserID, bookID int64) (*BookLoan, error)
	// Find by from user ID
	FindByFromUserID(ctx context.Context, fromUserID int64) ([]*BookLoan, error)
	// Find by to user ID
	FindByToUserID(ctx context.Context, fromUserID int64) ([]*BookLoan, error)
}
