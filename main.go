package main

import (
	"log"

	"github.com/1BarCode/go-bank-v1/app"
	"github.com/1BarCode/go-bank-v1/app/services"
	"github.com/1BarCode/go-bank-v1/db"
)


func main() {
	// load config / env variables

	// connect to db

	// create store
	store := db.NewStore()

	// create services
	services := services.NewServices(store)

	// create server
	server := app.NewServer(services)

	// start server
	err := server.Start(":8080")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}