package main

import (
	"database/sql"
	"log"

	"github.com/jerry-yue/simplebank/api"
	db "github.com/jerry-yue/simplebank/db/sqlc"
	"github.com/jerry-yue/simplebank/util"
	_ "github.com/lib/pq"
)

// viper insteaded!!!
// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:P@ssw0rd@127.0.0.1:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )
// viper insteaded!!!

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Load config file failed: ", err)
	}

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Connect to db failed: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Server started failed: ", err)
	}
}
