package api

import (
	db "github.com/Pokala15/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	transaction db.Transaction
	router      *gin.Engine
}

func NewServer(transaction db.Transaction) *Server {
	server := &Server{transaction: transaction}
	router := gin.Default()

	router.POST("/account/create", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.getAllAccounts)
	router.POST("/account/delete", server.deleteAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
