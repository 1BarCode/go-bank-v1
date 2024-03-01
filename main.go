package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/1BarCode/go-bank-v1/api"
	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	"github.com/1BarCode/go-bank-v1/services"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:admin123@localhost:5432/go_bank_v1?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	// load config / env variables

	// connect to db
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// create store with db connection
	store := db.NewStore(dbConn)

	// // create services
	services := services.NewServices(store)

	// // create server
	server := api.NewServer(services)

	// start server
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	fmt.Println("Server running on ", serverAddress)
}