package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/1BarCode/go-bank-v1/api"
	"github.com/1BarCode/go-bank-v1/config"
	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	"github.com/1BarCode/go-bank-v1/services"
	_ "github.com/lib/pq"
)



func main() {
	// load config / env variables
	env, err := config.LoadEnv(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect to db
	dbConn, err := sql.Open(env.DBDriver, env.DBSource)
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
	err = server.Start(env.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
	fmt.Println("Server running on ", env.ServerAddress)
}