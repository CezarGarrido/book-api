package entity

import (
	"context"
	"errors"
	"time"
)

var ErrBookNotFoud = errors.New("Não foi possível recuperar o livro")
var ErrBookCreate = errors.New("Não foi possível inserir o livro")

// Book represents the user's book structure
type Book struct {
	ID        int64     `json:"id"`      // Book id
	UserID    int64     `json:"user_id"` // User ID
	Title     string    `json:"title"`   // Book's title - required
	Pages     int       `json:"pages"`   // Number of pages - optional
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookUsecase interface {
	// Criar livro
	AddBookUserCollection(ctx context.Context, user User, book Book) (*Book, error)
	FindBookByID(ctx context.Context, bookID int64) (*Book, error)
	FindBooksByUserID(ctx context.Context, userID int64) ([]*Book, error)
}

type BookRepo interface {
	// Criar Livro no banco de dados
	Create(ctx context.Context, book *Book) (*Book, error)
	FindByID(ctx context.Context, bookID int64) (*Book, error)
	FindByUserID(ctx context.Context, userID int64) ([]*Book, error)
}
