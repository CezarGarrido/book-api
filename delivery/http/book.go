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

// AccountDeliveryHTTP
type BookDeliveryHTTP struct {
	bookUsecase entity.BookUsecase
	userUsecase entity.UserUsecase
}

func NewBookDeliveryHTTP(r *mux.Router, bookUsecase entity.BookUsecase, userUsecase entity.UserUsecase) {
	handler := &BookDeliveryHTTP{
		bookUsecase: bookUsecase,
		userUsecase: userUsecase,
	}
	r.HandleFunc("/book", handler.AddBookUserCollection).
		Name("add-book-collection").Methods("POST")
}

type NewBook struct {
	LoggedUserID int64 `json:"logged_user_id"`
	entity.Book
}

// AddBookUserCollection :
func (BookDelivery *BookDeliveryHTTP) AddBookUserCollection(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao ler json", http.StatusInternalServerError)
		return
	}

	var newBook NewBook

	err = json.Unmarshal(b, &newBook)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao decodificar json", http.StatusInternalServerError)
		return
	}

	user, err := BookDelivery.userUsecase.FindUserByID(ctx, newBook.LoggedUserID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookCreated, err := BookDelivery.bookUsecase.AddBookUserCollection(ctx, *user, entity.Book{
		Title:  newBook.Title,
		Pages:  newBook.Pages,
		UserID: newBook.LoggedUserID,
	})
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delivery.RespondWithJSON(w, bookCreated, http.StatusOK)
}
