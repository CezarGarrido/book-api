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

	db, err := infra.NewPostgres(infra.NewPostgresTestDSN())
	if err != nil {
		panic(err)
	}
	db.SQL.Exec(`DELETE FROM books`)
	db.SQL.Exec(`DELETE FROM book_loans`)
	db.SQL.Exec(`DELETE FROM users`)

	defer db.SQL.Exec(`DELETE FROM book_loans`)
	defer db.SQL.Exec(`DELETE FROM book_loans`)
	defer db.SQL.Exec(`DELETE FROM users`)

	userRepo := repository.NewUserPostgresRepo(db.SQL)

	userUsecase := usecase.NewUserUsecase(userRepo)

	userToCreate := newUser()

	userCreated, err := userUsecase.CreateUser(ctx, userToCreate)
	if err != nil {
		t.Fatal(err.Error())
	}

	AssertEqual(t, userCreated.Name, userToCreate.Name)
	AssertEqual(t, userCreated.Email, userToCreate.Email)
}

// Load user detais
func TestUsecaseLoadUserDetails(t *testing.T) {
	ctx := context.TODO()

	db, err := infra.NewPostgres(infra.NewPostgresTestDSN())
	if err != nil {
		panic(err)
	}
	db.SQL.Exec(`DELETE FROM books`)
	db.SQL.Exec(`DELETE FROM book_loans`)
	db.SQL.Exec(`DELETE FROM users`)

	defer db.SQL.Exec(`DELETE FROM users`)

	userRepo := repository.NewUserPostgresRepo(db.SQL)
	bookRepo := repository.NewBookPostgresRepo(db.SQL)

	userUsecase := usecase.NewUserUsecase(userRepo)
	bookUsecase := usecase.NewBookUsecase(bookRepo)

	userToCreate := newUser()

	user, err := userUsecase.CreateUser(ctx, userToCreate)
	if err != nil {
		t.Fatal(err.Error())
	}

	AssertEqual(t, user.Name, userToCreate.Name)
	AssertEqual(t, user.Email, userToCreate.Email)

	findUser, err := userUsecase.FindUserByID(ctx, user.ID)
	if err != nil {
		t.Fatal(err.Error())
	}

	collection, err := bookUsecase.FindBooksByUserID(ctx, findUser.ID)
	if err != nil {
		t.Fatal(err.Error())
	}

	lenCol := len(collection)

	if lenCol > 0 {
		t.Fatal("Expected collection empty")
	}
}

// Add new book collection
func TestUsecaseAddBookUserCollection(t *testing.T) {
	ctx := context.TODO()

	db, err := infra.NewPostgres(infra.NewPostgresTestDSN())
	if err != nil {
		panic(err)
	}
	db.SQL.Exec(`DELETE FROM books`)
	db.SQL.Exec(`DELETE FROM book_loans`)
	db.SQL.Exec(`DELETE FROM users`)

	defer db.SQL.Exec(`DELETE FROM books`)
	defer db.SQL.Exec(`DELETE FROM users`)

	bookRepo := repository.NewBookPostgresRepo(db.SQL)
	userRepo := repository.NewUserPostgresRepo(db.SQL)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userToCreate := newUser()
	user, err := userUsecase.CreateUser(ctx, userToCreate)
	if err != nil {
		t.Fatal(err.Error())
	}
	AssertEqual(t, user.Name, userToCreate.Name)
	AssertEqual(t, user.Email, userToCreate.Email)

	bookUsecase := usecase.NewBookUsecase(bookRepo)

	bookToCreate := newBook()

	bookCreated, err := bookUsecase.AddBookUserCollection(ctx, *user, bookToCreate)

	if err != nil {
		t.Fatal(err.Error())
	}

	AssertEqual(t, bookToCreate.Title, bookCreated.Title)
	AssertEqual(t, bookToCreate.Pages, bookCreated.Pages)
}

// Add new book collection
func TestUsecaseNewLoan(t *testing.T) {
	ctx := context.TODO()

	db, err := infra.NewPostgres(infra.NewPostgresTestDSN())
	if err != nil {
		panic(err)
	}
	db.SQL.Exec(`DELETE FROM books`)
	db.SQL.Exec(`DELETE FROM book_loans`)
	db.SQL.Exec(`DELETE FROM users`)

	defer db.SQL.Exec(`DELETE FROM book_loans`)
	defer db.SQL.Exec(`DELETE FROM users`)

	bookRepo := repository.NewBookPostgresRepo(db.SQL)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	userRepo := repository.NewUserPostgresRepo(db.SQL)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userToCreate := newUser()
	user, err := userUsecase.CreateUser(ctx, userToCreate)
	if err != nil {
		t.Fatal(err.Error())
	}
	AssertEqual(t, user.Name, userToCreate.Name)
	AssertEqual(t, user.Email, userToCreate.Email)

	bookToCreate := newBook()

	bookCreated, err := bookUsecase.AddBookUserCollection(ctx, *user, bookToCreate)
	if err != nil {
		t.Fatal(err.Error())
	}

	AssertEqual(t, bookToCreate.Title, bookCreated.Title)
	AssertEqual(t, bookToCreate.Pages, bookCreated.Pages)

	userToCreat2 := entity.User{
		Name:  "Teste2",
		Email: "teste2@mail.com",
	}

	userCreated, err := userUsecase.CreateUser(ctx, userToCreat2)
	if err != nil {
		t.Fatal(err.Error())
	}
	AssertEqual(t, userCreated.Name, userToCreat2.Name)
	AssertEqual(t, userCreated.Email, userToCreat2.Email)

	bookLoanRepo := repository.NewBookLoanPostgresRepo(db.SQL)

	bookLoanUsecase := usecase.NewBookLoanUsecase(bookLoanRepo)

	bookLoanCreated, err := bookLoanUsecase.LendBook(ctx, *user, userCreated.ID, bookCreated.ID)
	if err != nil {
		t.Fatal(err.Error())
	}

	AssertEqual(t, bookLoanCreated.BookID, bookCreated.ID)
	AssertEqual(t, bookLoanCreated.ToUserID, userCreated.ID)
	AssertEqual(t, bookLoanCreated.FromUserID, user.ID)

}

// Fail test duplicate email
func TestUsecaseFailCreateUser(t *testing.T) {
	ctx := context.TODO()

	db, err := infra.NewPostgres(infra.NewPostgresTestDSN())
	if err != nil {
		panic(err)
	}
	db.SQL.Exec(`DELETE FROM books`)
	db.SQL.Exec(`DELETE FROM book_loans`)
	db.SQL.Exec(`DELETE FROM users`)
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

func newBook() entity.Book {
	return entity.Book{
		Title: "Teste",
		Pages: 10,
	}
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}
