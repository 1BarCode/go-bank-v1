# start the postgres container
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin123 -d postgres:12-alpine

# create the database named go_bank_v1 - then connect to the database (not root db)
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root go_bank_v1

dropdb:
	docker exec -it postgres12 dropdb go_bank_v1

# path to migration folder
migrateup:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/go_bank_v1?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/go_bank_v1?sslmode=disable" -verbose down	

# generate the db docs
db_docs:
	dbdocs build doc/db.dbml

# generate the db schema from the .dbml file
db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml


migratedown:
.PHONY: postgres createdb dropdb migrateup db_docs db_schema