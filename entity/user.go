package entity

import (
	"context"
	"errors"
	"time"
)

var ErrDuplicatedEmail = errors.New("Email já cadastrado")
var ErrUserNotFoud = errors.New("Não foi possível recuperar o usuário")

// Represents the system user
type User struct {
	ID            int64       `json:"id"`             // User id
	Name          string      `json:"name"`           // User name - required
	Email         string      `json:"email"`          // User email - required - unique
	Collection    []*Book     `json:"collection"`     // All user books
	LentBooks     []*BookLoan `json:"lent_books"`     // All books that were loaned to other users
	BorrowedBooks []*BookLoan `json:"borrowed_books"` // All books that were borrowed from other users
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}
type UserUsecase interface {
	CreateUser(ctx context.Context, user User) (*User, error)
	FindAllUsers(ctx context.Context) ([]*User, error)
	FindUserByID(ctx context.Context, id int64) (*User, error)
}

type UserRepo interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindAllUsers(ctx context.Context) ([]*User, error)
	FindUserByID(ctx context.Context, id int64) (*User, error)
}
