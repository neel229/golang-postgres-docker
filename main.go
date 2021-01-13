package main

import (
	"database/sql"
	"log"

	"github.com/neel229/golang-postgres-docker/api"
	db "github.com/neel229/golang-postgres-docker/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:postgres@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the Database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.StartServer(serverAddress)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
