package usecase

import (
	"context"

	"github.com/CezarGarrido/book-api/entity"
)

type userUsecase struct {
	userRepo entity.UserRepo
}

func NewAccountUsecase(userRepo entity.UserRepo) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (userUsecase *userUsecase) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	return userUsecase.userRepo.Create(ctx, user)
}
