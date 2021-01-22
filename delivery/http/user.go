package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

func (userDelivery *UserDeliveryHTTP) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var newUser entity.User

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error parse body", err.Error())
		return
	}

	err = json.Unmarshal(b, &newUser)
	if err != nil {
		log.Println("Error json unmarshal new account:", err.Error())
		return
	}

	userDelivery.userUsecase.CreateUser(ctx, newUser)
}
