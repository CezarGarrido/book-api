package entity

import (
	"context"
	"errors"
	"time"
)

var ErrDuplicatedEmail = errors.New("Email já cadastrado")
var ErrUserNotFoud = errors.New("Não foi possível recuperar o usuário")

type User struct {
	ID            int64       `json:"id"`
	Name          string      `json:"name"`
	Email         string      `json:"email"`
	Password      string      `json:"password"`
	Collection    []*Book     `json:"collection"`
	LentBooks     []*BookLoan `json:"lent_books"`
	BorrowedBooks []*BookLoan `json:"borrowed_books"`
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
