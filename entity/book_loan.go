package entity

import (
	"context"
	"time"
)

type BookLoan struct {
	BookID     int64
	FromUserID int64
	ToUserID   int64
	LentAt     time.Time
	ReturnedAt *time.Time
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
