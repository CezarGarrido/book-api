package main

import (
	"log"

	"github.com/CezarGarrido/book-api/infra"
)

func main() {
	log.Println("Running api")

	db, err := infra.NewPostgres(infra.NewPostgresDSN())
	if err != nil {
		panic(err)
	}

	err = infra.NewMigratePostgres(db.SQL)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("All migrations Ok")
}
