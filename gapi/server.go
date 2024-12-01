package gapi

import (
	"fmt"
	db "simple_bank/db/sqlc"
	"simple_bank/pb"
	"simple_bank/token"
	"simple_bank/util"
	"simple_bank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token master: %w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker, taskDistributor: taskDistributor}

	return server, nil
}
