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
	router := mux.NewRouter()
	db, err := infra.NewPostgres(infra.NewPostgresDSN())
	if err != nil {
		panic(err)
	}
	err = infra.NewMigratePostgres(db.SQL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	dbTests, err := infra.NewPostgres(infra.NewPostgresTestDSN())
	if err != nil {
		panic(err)
	}

	err = infra.NewMigratePostgres(dbTests.SQL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	userRepo := repository.NewUserPostgresRepo(db.SQL)
	bookRepo := repository.NewBookPostgresRepo(db.SQL)
	bookLoanRepo := repository.NewBookLoanPostgresRepo(db.SQL)

	userUsecase := usecase.NewUserUsecase(userRepo)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookLoanUsecase := usecase.NewBookLoanUsecase(bookLoanRepo)

	deliveryHTTP.NewUserDeliveryHTTP(router, userUsecase, bookUsecase, bookLoanUsecase)
	deliveryHTTP.NewBookLoanDeliveryHTTP(router, bookLoanUsecase, bookUsecase, userUsecase)
	deliveryHTTP.NewBookDeliveryHTTP(router, bookUsecase, userUsecase)

	log.Println("Server running on port :8089")

	log.Fatal(http.ListenAndServe(":8089", router))
}
