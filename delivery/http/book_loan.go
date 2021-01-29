package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CezarGarrido/book-api/delivery"
	"github.com/CezarGarrido/book-api/entity"
	"github.com/gorilla/mux"
)

// BookLoanDeliveryHTTP
type BookLoanDeliveryHTTP struct {
	bookLoanUsecase entity.BookLoanUsecase
	bookUsecase     entity.BookUsecase
	userUsecase     entity.UserUsecase
}

func NewBookLoanDeliveryHTTP(r *mux.Router, bookLoanUsecase entity.BookLoanUsecase, bookUsecase entity.BookUsecase, userUsecase entity.UserUsecase) {
	handler := &BookLoanDeliveryHTTP{
		bookLoanUsecase: bookLoanUsecase,
		bookUsecase:     bookUsecase,
		userUsecase:     userUsecase,
	}
	r.HandleFunc("/book/lend", handler.LendBook).
		Name("create-loan").Methods("POST")

	r.HandleFunc("/book/lend/return", handler.ReturnBook).
		Name("return-book-loan").Methods("PUT")
}

type NewBookLoan struct {
	LoggedUserID int64 `json:"logged_user_id"`
	entity.BookLoan
}

// Lend book
func (bookLoanDelivery *BookLoanDeliveryHTTP) LendBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao ler json", http.StatusInternalServerError)
		return
	}

	var newBookLoan NewBookLoan

	err = json.Unmarshal(b, &newBookLoan)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao decodificar json", http.StatusInternalServerError)
		return
	}

	user, err := bookLoanDelivery.userUsecase.FindUserByID(ctx, newBookLoan.LoggedUserID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toUser, err := bookLoanDelivery.userUsecase.FindUserByID(ctx, newBookLoan.ToUserID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book, err := bookLoanDelivery.bookUsecase.FindBookByID(ctx, newBookLoan.BookID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookLoanCreated, err := bookLoanDelivery.bookLoanUsecase.LendBook(ctx, *user, toUser.ID, book.ID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delivery.RespondWithJSON(w, bookLoanCreated, http.StatusOK)
}

// Return Book
func (bookLoanDelivery *BookLoanDeliveryHTTP) ReturnBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao ler json", http.StatusInternalServerError)
		return
	}

	var newBookLoan NewBookLoan

	err = json.Unmarshal(b, &newBookLoan)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao decodificar json", http.StatusInternalServerError)
		return
	}
	user, err := bookLoanDelivery.userUsecase.FindUserByID(ctx, newBookLoan.LoggedUserID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book, err := bookLoanDelivery.bookUsecase.FindBookByID(ctx, newBookLoan.BookID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookLoanReturned, err := bookLoanDelivery.bookLoanUsecase.ReturnBook(ctx, *user, book.ID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delivery.RespondWithJSON(w, bookLoanReturned, http.StatusOK)
}
