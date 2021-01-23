package entity

import (
	"context"
	"time"
)

type BookLoan struct {
	ID         int64      `json:"id"`
	BookID     int64      `json:"book_id"`
	FromUserID int64      `json:"from_user_id"`
	ToUserID   int64      `json:"to_user_id"`
	LentAt     time.Time  `json:"lent_at"`
	ReturnedAt *time.Time `json:"returned_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type BookLoanUsecase interface {
	// Emprestar o livro
	LendBook(user User, toUserID int64, bookID int64) (BookLoan, error)
	// Devolver o livro
	ReturnBook(user User, bookID int64) (BookLoan, error)
}

type BookLoanRepo interface {
	// Criar o livro no banco de dados
	Create(ctx context.Context, bookLoan *BookLoan) (*BookLoan, error)
	// Atualiza o empréstimo como devolvido
	// Atualiza a data da devolução
	ReturnBook(ctx context.Context, bookID int64) (*BookLoan, error)
}
