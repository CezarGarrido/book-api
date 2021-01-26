package tests

import (
	"context"
	"reflect"
	"testing"

	"github.com/CezarGarrido/book-api/entity"
	"github.com/CezarGarrido/book-api/infra"
	"github.com/CezarGarrido/book-api/repository"
	"github.com/CezarGarrido/book-api/usecase"
)

func TestUsecaseCreateUser(t *testing.T) {
	ctx := context.TODO()

	db, err := infra.NewPostgres(infra.NewPostgresDSN())
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserPostgresRepo(db.SQL)

	userUsecase := usecase.NewUserUsecase(userRepo)

	userToCreate := newUser()

	userCreated, err := userUsecase.CreateUser(ctx, userToCreate)
	if err != nil {
		t.Fatal(err.Error())
	}

	AssertEqual(t, userCreated.Name, userToCreate.Name)
	AssertEqual(t, userCreated.Email, userToCreate.Email)
	AssertEqual(t, userCreated.Password, userToCreate.Password)
}

func newUser() entity.User {
	return entity.User{
		Name:     "User test",
		Email:    "user@mail.com",
		Password: "102030@@",
	}
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}
