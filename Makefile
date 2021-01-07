postgres:
	docker run --name demo -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it demo createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it demo dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@172.17.212.23:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@172.17.212.23:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
