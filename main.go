package main

import (
	"database/sql"
	"log"

	"github.com/neel229/golang-postgres-docker/api"
	db "github.com/neel229/golang-postgres-docker/db/sqlc"
	"github.com/neel229/golang-postgres-docker/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the Database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.StartServer(config.ServerAddress)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
