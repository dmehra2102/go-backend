package gapi

import (
	"fmt"
	db "simple_bank/db/sqlc"
	"simple_bank/pb"
	"simple_bank/token"
	"simple_bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.SymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token master: %w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	return server, nil
}
