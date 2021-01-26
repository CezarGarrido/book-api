package usecase

import (
	"context"

	"github.com/CezarGarrido/book-api/entity"
)

type userUsecase struct {
	userRepo entity.UserRepo
	bookLoanRepo entity.BookLoanRepo
	bookRepo entity.BookRepo
}

func NewUserUsecase(userRepo entity.UserRepo) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (userUsecase *userUsecase) CreateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	return userUsecase.userRepo.Create(ctx, &user)
}

func (userUsecase *userUsecase) FindAllUsers(ctx context.Context) ([]*entity.User, error) {
	return nil, nil
}

func (userUsecase *userUsecase) FindUserByID(ctx context.Context, id int64) (*entity.User, error) {
	return userUsecase.userRepo.FindUserByID(ctx, id)
}
