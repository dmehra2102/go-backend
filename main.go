package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"simple_bank/api"
	db "simple_bank/db/sqlc"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to DB: %v", err)
	}

	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	go func() {
		if err = server.Start(serverAddress); err != nil {
			log.Fatalf("Cannot start server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down server...")
	cancel()

}
