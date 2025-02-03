package main

import (
	"database/sql"
	"log"

	"github.com/Pokala15/simplebank/api"
	db "github.com/Pokala15/simplebank/db/sqlc"
	"github.com/Pokala15/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("couldn't load config file: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	transaction := db.NewTransaction(conn)
	server := api.NewServer(transaction)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
