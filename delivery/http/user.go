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

// AccountDeliveryHTTP
type UserDeliveryHTTP struct {
	userUsecase     entity.UserUsecase
	bookUsecase     entity.BookUsecase
	bookLoanUsecase entity.BookLoanUsecase
}

func NewUserDeliveryHTTP(r *mux.Router, userUsecase entity.UserUsecase, bookUsecase entity.BookUsecase, bookLoanUsecase entity.BookLoanUsecase) {
	handler := &UserDeliveryHTTP{
		userUsecase:     userUsecase,
		bookUsecase:     bookUsecase,
		bookLoanUsecase: bookLoanUsecase,
	}
	r.HandleFunc("/users", handler.Create).
		Name("create-user").Methods("POST")
	r.HandleFunc("/users/{user_id:[0-9]+}", handler.FindUserByID).
		Name("create-user").Methods("GET")
}

// Create :
func (userDelivery *UserDeliveryHTTP) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var newUser entity.User

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &newUser)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userCreated, err := userDelivery.userUsecase.CreateUser(ctx, newUser)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delivery.RespondWithJSON(w, userCreated, http.StatusOK)
}

// FindUserByID :
func (userDelivery *UserDeliveryHTTP) FindUserByID(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userID, _ := strconv.ParseInt(mux.Vars(r)["user_id"], 10, 64)

	user, err := userDelivery.userUsecase.FindUserByID(ctx, userID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lentBooks, err := userDelivery.bookLoanUsecase.FindByFromUserID(ctx, user.ID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.LentBooks = lentBooks

	borrowedBooks, err := userDelivery.bookLoanUsecase.FindByToUserID(ctx, user.ID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.BorrowedBooks = borrowedBooks

	collection, err := userDelivery.bookUsecase.FindBooksByUserID(ctx, user.ID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Collection = collection

	delivery.RespondWithJSON(w, user, http.StatusOK)
}
