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

var user *entity.User

func TestUsecaseCreateUser(t *testing.T) {
	ctx := context.TODO()

	db, err := infra.NewPostgres(infra.NewPostgresTestDSN())
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
	user = userCreated
}

func TestUsecaseLoadUserDetails(t *testing.T) {
	ctx := context.TODO()

	db, err := infra.NewPostgres(infra.NewPostgresTestDSN())
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserPostgresRepo(db.SQL)
	bookRepo := repository.NewBookPostgresRepo(db.SQL)

	userUsecase := usecase.NewUserUsecase(userRepo)
	bookUsecase := usecase.NewBookUsecase(bookRepo)

	user, err := userUsecase.FindUserByID(ctx, user.ID)
	if err != nil {
		t.Fatal(err.Error())
	}

	collection, err := bookUsecase.FindBooksByUserID(ctx, user.ID)
	if err != nil {
		t.Fatal(err.Error())
	}

	lenCol := len(collection)

	if lenCol > 0 {
		t.Fatal("Expected collection empty")
	}
}

// Fail test duplicate email
func TestUsecaseFailCreateUser(t *testing.T) {
	ctx := context.TODO()

	db, err := infra.NewPostgres(infra.NewPostgresTestDSN())
	if err != nil {
		panic(err)
	}

	defer db.SQL.Exec(`DELETE FROM users`)

	userRepo := repository.NewUserPostgresRepo(db.SQL)

	userUsecase := usecase.NewUserUsecase(userRepo)

	userToCreate := newUser()

	_, err = userUsecase.CreateUser(ctx, userToCreate)

	if err != nil {
		expected := "Email já cadastrado"
		AssertEqual(t, err.Error(), expected)
	}

}

func newUser() entity.User {
	return entity.User{
		Name:  "User test",
		Email: "user@mail.com",
	}
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}
