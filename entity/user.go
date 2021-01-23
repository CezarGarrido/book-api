package entity

import (
	"context"
	"errors"
	"strings"
	"time"
)

var ErrNameRequired = errors.New("\"Name\" length must be at least 8 characters long")
var ErrEmailRequired = errors.New("\"email\" is required")
var ErrEmailInvalid = errors.New("\"email\" must be a valid email")
var ErrPasswordInvalid = errors.New("\"password\" length must be 6 characters long")
var ErrPasswordRequired = errors.New("\"password\" is required")
var ErrUserIsExists = errors.New("Usuário já existe")
var ErrDuplicatedEmail = errors.New("Email já cadastrado.")

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

func (user *User) Validate() error {
	if len(strings.Trim(user.Name, " ")) < 8 {
		return ErrNameRequired
	}

	return nil
}

type UserUsecase interface {
	CreateUser(ctx context.Context, user User) (*User, error)
	FindAllUsers(ctx context.Context) ([]*User, error)
}

type UserRepo interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindAllUsers(ctx context.Context) ([]*User, error)
}
