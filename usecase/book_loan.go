package usecase

import (
	"context"
	"time"

	"github.com/CezarGarrido/book-api/entity"
)

type bookLoanUsecase struct {
	bookLoanRepo entity.BookLoanRepo
}

func NewBookLoanUsecase(bookLoanRepo entity.BookLoanRepo) *bookLoanUsecase {
	return &bookLoanUsecase{
		bookLoanRepo: bookLoanRepo,
	}
}

// Emprestar o livro
func (bookLoanUsecase *bookLoanUsecase) LendBook(ctx context.Context, user entity.User, toUserID int64, bookID int64) (*entity.BookLoan, error) {
	return bookLoanUsecase.bookLoanRepo.Create(ctx, &entity.BookLoan{
		FromUserID: user.ID,
		ToUserID:   toUserID,
		BookID:     bookID,
		LentAt:     time.Now(),
	})
}

// Devolver o livro
func (bookLoanUsecase *bookLoanUsecase) ReturnBook(ctx context.Context, user entity.User, bookID int64) (*entity.BookLoan, error) {
	bookLoan, err := bookLoanUsecase.bookLoanRepo.FindByFromUserAndBookID(ctx, user.ID, bookID)
	if err != nil {
		return nil, entity.ErrBookLoanNotFoud
	}

	if bookLoan.ReturnedAt != nil {
		return nil, entity.ErrBookLoanIsReturned
	}

	now := time.Now()
	bookLoan.ReturnedAt = &now
	return bookLoanUsecase.bookLoanRepo.ReturnBook(ctx, bookLoan)
}

// Devolver o livro
func (bookLoanUsecase *bookLoanUsecase) FindByToUserID(ctx context.Context, toUserID int64) ([]*entity.BookLoan, error) {

	return bookLoanUsecase.bookLoanRepo.FindByToUserID(ctx, toUserID)
}

func (bookLoanUsecase *bookLoanUsecase) FindByFromUserID(ctx context.Context, toUserID int64) ([]*entity.BookLoan, error) {

	return bookLoanUsecase.bookLoanRepo.FindByFromUserID(ctx, toUserID)
}
