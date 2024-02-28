# Go Bank 1.0

## Description

This is a simple bank application that allows users to create an account, deposit, withdraw, and transfer money between accounts.

## Technologies

- Golang/Gin
- PostgreSQL
- Docker
- Docker Compose
- Makefile
- Sqlc
- Gomock
- Testify
- Golang-migrate

## Features

## Installation

<!-- 1. Clone the repository
2. Run `make compose-up` to start the application
3. Run `make compose-down` to stop the application
4. Run `make compose-logs` to view the logs of the application
5. Run `make compose-psql` to connect to the database
6. Run `make compose-test` to run the tests
7. Run `make compose-migrate-up` to run the migrations
8. Run `make compose-migrate-down` to rollback the migrations -->

## Usage

- Gin server is running on port 8080. Start the server by running `run go main.go` in the root directory of the project.
- The database is running on port 5432 in the docker container.

     - If connecting to the DB for the first time:
          1. Run `make postgresinit` to initialize the docker container containing the database.
          2. Run `make createdb` to create the database.
          3. Connect to the DB with TablePlus or any other DB client.
          4. Run `make migrateup` to run the migrations.

- SQLC is used to generate idiomatic Go code for working with the database based on the SQL queries in the `db/query` folder.
     - To generate the SQLC code, run `make sqlc` in the root directory of the project.
     - SQLC looks at the sqlc.yaml file to generate the code.
     - The generated code is placed in the `db/sqlc` directory.
          - `models.go` contains the structs for the tables in the database.
          - `db.go` contains the Queries struct for using with a sql.DB object or with a sql.Tx object for transactions.
          - The Go code ends with `.sql.go`.
