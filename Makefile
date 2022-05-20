postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=P@ssw0rd -p 5432:5432 -d postgres:12.10-alpine 

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:P@ssw0rd@127.0.0.1:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:P@ssw0rd@127.0.0.1:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migratedown migrateup sqlc test
