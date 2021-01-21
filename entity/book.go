package entity

import "time"

type Book struct {
	ID          int64     `json:"id"`
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	AuthorName  string    `json:"author_name"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookUsecase interface {
	Create(book Book) (error, Book)
}
