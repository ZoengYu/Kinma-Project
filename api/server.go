package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kinmaBackend/db/sqlc"
)

type Server struct {
	store 		*db.Store
	router 		*gin.Engine
}

func NewServer(store *db.Store) *Server{
	server := &Server{store : store}
	router := gin.Default()

	//Restful API generated here
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/products", server.createProduct)
	router.GET("/products/:id", server.getProduct)
	router.GET("/products", server.listProduct)
	router.PUT("/updateproduct/:id", server.updateProduct)

	router.POST("/fundraise", server.createFundraise)
	router.GET("/fundraise/:id", server.getFundraise)
	router.PUT("/exitfundraise", server.exitFundraise)
	router.PUT("/addfundaiseamount", server.addFundraise)
	
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}