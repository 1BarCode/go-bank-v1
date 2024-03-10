package services

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/1BarCode/go-bank-v1/config"
	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	_ "github.com/lib/pq"
)

// main test file for db package

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	env, err := config.LoadEnv("..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(env.DBDriver, env.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}