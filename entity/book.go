package entity

import (
	"context"
	"time"
)

type Book struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Pages       int       `json:"pages"`
	AuthorName  string    `json:"author_name"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookUsecase interface {
	// Criar livro
	CreateBook(user User, book Book) (Book, error)
}

type BookRepo interface {
	// 
	Create(ctx context.Context, book *Book) (*Book, error)
}
