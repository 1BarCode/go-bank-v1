# start the postgres container
pginit:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin123 -d postgres:12-alpine

pgstart:
	docker start postgres12

pgstop:
	docker stop postgres12

# create the database named go_bank_v1 - then connect to the database (not root db)
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root go_bank_v1

dropdb:
	docker exec -it postgres12 dropdb go_bank_v1

# create a new migration file with custom name
# ex: make new_migration name=create_users_table
new_migration:
	migrate create -ext sql -dir db/migration -seq "$(name)"

# path to migration folder
migrateup:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/go_bank_v1?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/go_bank_v1?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/go_bank_v1?sslmode=disable" -verbose down	

migratedown1:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/go_bank_v1?sslmode=disable" -verbose down 1

# generate the db docs - there will be a link to the db docs in the terminal
db_docs:
	dbdocs build doc/db.dbml

# generate the db schema from the .dbml file
db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

# run unit tests in all subdirectories
test:
	go test -v -cover ./...

server:
	go run main.go

# generate a mock for the Store interface and save it in the db/mock folder
mockdb:
	mockgen -package mockdb -destination db/mock/store.go github.com/1BarCode/go-bank-v1/db/sqlc Store

mockservice:
	mockgen -package mockservices -destination services/mock/services.go github.com/1BarCode/go-bank-v1/services Services

.PHONY: pginit pgstart pgstop createdb dropdb migrateup migrateup1 migratedown migratedown1 db_docs db_schema sqlc server mockdb mockservice