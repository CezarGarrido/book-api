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
type UserDeliveryHTTP struct {
	userUsecase entity.UserUsecase
}

func NewUserDeliveryHTTP(r *mux.Router, userUsecase entity.UserUsecase) {
	handler := &UserDeliveryHTTP{
		userUsecase: userUsecase,
	}
	r.HandleFunc("/users", handler.Create).
		Name("create-user").Methods("POST")
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
