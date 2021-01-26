package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
	r.HandleFunc("/users/{user_id:[0-9]+}/books", handler.AddBookUserCollection).
		Name("add-book-collection").Methods("POST")
}

// AddBookUserCollection :
func (BookDelivery *BookDeliveryHTTP) AddBookUserCollection(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userID, _ := strconv.ParseInt(mux.Vars(r)["user_id"], 10, 64)
	log.Println(userID)

	user, err := BookDelivery.userUsecase.FindUserByID(ctx, userID)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(user)
	
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao ler json", http.StatusInternalServerError)
		return
	}

	var newBook entity.Book

	err = json.Unmarshal(b, &newBook)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, "Erro ao decodificar json", http.StatusInternalServerError)
		return
	}

	bookCreated, err := BookDelivery.bookUsecase.AddBookUserCollection(ctx, *user, newBook)
	if err != nil {
		log.Println(err.Error())
		delivery.RespondWithJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delivery.RespondWithJSON(w, bookCreated, http.StatusOK)
}
