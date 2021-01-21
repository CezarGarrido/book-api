package main

import (
	"log"

	"github.com/CezarGarrido/book-api/infra"
)


func main() {
	log.Println("Running api")

	db, err := infra.NewPostgres("")
	if err != nil {
		panic(err)
	}

	err = infra.NewMigrate(db.SQL)
	if err != nil {
		panic(err)
	}

}
