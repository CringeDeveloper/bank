package main

import (
	"database/sql"
	"github.com/cringedeveloper/bank/api"
	db "github.com/cringedeveloper/bank/db/sqlc"
	"github.com/cringedeveloper/bank/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read config")
	}

	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
