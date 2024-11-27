package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"simple_bank/api"
	db "simple_bank/db/sqlc"
	"simple_bank/util"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to DB: %v", err)
	}

	defer conn.Close()

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	go func() {
		if err = server.Start(config.ServerAddress); err != nil {
			log.Fatalf("Cannot start server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down server...")
	cancel()

}
