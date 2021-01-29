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
	// Add book in user collection
	AddBookUserCollection(ctx context.Context, user User, book Book) (*Book, error)
	// Return book by ID
	FindBookByID(ctx context.Context, bookID int64) (*Book, error)
	// Find all user books
	FindBooksByUserID(ctx context.Context, userID int64) ([]*Book, error)
}

type BookRepo interface {
	// Create book in database
	Create(ctx context.Context, book *Book) (*Book, error)
	// Return book by ID
	FindByID(ctx context.Context, bookID int64) (*Book, error)
	// Find all user books
	FindByUserID(ctx context.Context, userID int64) ([]*Book, error)
}
