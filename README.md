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

- The database is running on port 5432 in the docker container.

     - If connecting to the DB for the first time:
          1. Run `make pgstart` to initialize the docker database container.
          2. Run `make createdb` to create the database.
          3. Connect to the DB with TablePlus or any other DB client.
          4. Run `make migrateup` to run the migrations.
     - For subsequent connections:
          1. Run `make postgresstart` to start the database container.
          2. To stop the database container, run `make pgstop`.

- Gin server is running on port 8080. Start the server by running `make server` in the root directory of the project. Make sure the database is running before starting the server.

- SQLC is used to generate idiomatic Go code for working with the database based on the pre-defined SQL queries we've written in the `db/query` folder.

     - SQLC looks at the `sqlc.yaml` file in the root directory to configure and generate the code.
     - To add new/update SQL queries:
          1. Write the queries in the `db/query` folder.
          2. To generate the SQLC code, run `make sqlc` in the root directory of the project.
          3. The generated code is placed in the `db/sqlc` directory.
                - `models.go` contains the structs for the tables in the database.
                - `db.go` contains the Queries struct for using with a sql.DB object or with a sql.Tx object for transactions.
                - The Go code ends with `.sql.go`.
          4. Make sure to create the corresponding `<table>_test.go` file in the `db/sqlc` directory to test the generated code.

- Gomock / Mockgen is used to generate a mock db and mock services for testing the 2 layers.
     1. To generate the mock db, run `make mockdb` in the root directory of the project.
     2. To generate the mock services, run `make mockservice` in the root directory of the project.
