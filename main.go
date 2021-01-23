package main

import (
	"log"
	"net/http"

	deliveryHTTP "github.com/CezarGarrido/book-api/delivery/http"
	"github.com/CezarGarrido/book-api/infra"
	"github.com/CezarGarrido/book-api/repository"
	"github.com/CezarGarrido/book-api/usecase"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting API")

	router := mux.NewRouter()

	db, err := infra.NewPostgres(infra.NewPostgresDSN())
	if err != nil {
		panic(err)
	}
	log.Println("Running migrations")
	err = infra.NewMigratePostgres(db.SQL)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Successful migrations")

	userRepo := repository.NewUserPostgresRepo(db.SQL)

	userUsecase := usecase.NewUserUsecase(userRepo)

	deliveryHTTP.NewUserDeliveryHTTP(router, userUsecase)

	log.Println("Server running on port :8089")

	log.Fatal(http.ListenAndServe(":8089", router))
}
