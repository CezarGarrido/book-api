package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
	r.HandleFunc("/users/{user_id:[0-9]+}/books/loans", handler.LendBook).
		Name("create-loan").Methods("POST")

	r.HandleFunc("/users/{user_id:[0-9]+}/books/loans/return", handler.ReturnBook).
		Name("return-book-loan").Methods("PUT")
}

// Emprestar o livro
func (bookLoanDelivery *BookLoanDeliveryHTTP) LendBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, _ := strconv.ParseInt(mux.Vars(r)["user_id"], 10, 64)
	user, err := bookLoanDelivery.userUsecase.FindUserByID(ctx, userID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao ler json", http.StatusInternalServerError)
		return
	}

	var newBookLoan entity.BookLoan

	err = json.Unmarshal(b, &newBookLoan)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao decodificar json", http.StatusInternalServerError)
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

// Devolver o livro
func (bookLoanDelivery *BookLoanDeliveryHTTP) ReturnBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, _ := strconv.ParseInt(mux.Vars(r)["user_id"], 10, 64)
	user, err := bookLoanDelivery.userUsecase.FindUserByID(ctx, userID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao ler json", http.StatusInternalServerError)
		return
	}

	var newBookLoan entity.BookLoan

	err = json.Unmarshal(b, &newBookLoan)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao decodificar json", http.StatusInternalServerError)
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
