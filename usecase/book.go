package usecase

import (
	"context"

	"github.com/CezarGarrido/book-api/entity"
)

type bookUsecase struct {
	bookRepo entity.BookRepo
}

func NewBookUsecase(bookRepo entity.BookRepo) *bookUsecase {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}

func (bookUsecase *bookUsecase) AddBookUserCollection(ctx context.Context, user entity.User, book entity.Book) (*entity.Book, error) {
	book.UserID = user.ID
	return bookUsecase.bookRepo.Create(ctx, &book)
}

func (bookUsecase *bookUsecase) FindBookByID(ctx context.Context, bookID int64) (*entity.Book, error) {
	return bookUsecase.bookRepo.FindByID(ctx, bookID)
}
