package usecase

import (
	"context"

	"github.com/CezarGarrido/book-api/entity"
)

type userUsecase struct {
	userRepo entity.UserRepo
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
