package api

import (
	"fmt"
	db "simple_bank/db/sqlc"
	"simple_bank/token"
	"simple_bank/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.SymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token master: %w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// Adding Routes
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account/all", server.listAccounts)
	router.PATCH("/account/:id", server.updateAccount)
	router.DELETE("/account/:id", server.deleteAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server, nil
}

// This function runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
