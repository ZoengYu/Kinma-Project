package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/kinmaBackend/db/sqlc"
	"github.com/kinmaBackend/token"
	"github.com/kinmaBackend/util"
)

type Server struct {
	store 			*db.Store
	tokenMaker 	token.Maker
	router 			*gin.Engine
	config      util.Config
}

func NewServer(config util.Config, store *db.Store) (*Server, error){
	//here we use AsymmetricPaseto token
	tokenMaker, err := token.NewAsymmetricPasetoMaker()
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store 		: store,
		config		: config,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency",validCurrency)
	}

	server.setupRouter()
	
	return server, nil
}

func (server *Server) setupRouter() {
		router := gin.Default()
		//Restful API generated here
		router.POST("/users", server.createUser)
		router.POST("/users/login", server.loginUser)
		
		authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

		authRoutes.POST("/myaccounts", server.createAccount)
		authRoutes.GET("/myaccounts/:id", server.getAccount)
		authRoutes.GET("/myaccounts", server.listMyAccount)

		authRoutes.GET("/myproducts/:product_id", server.getMyProduct)
		authRoutes.GET("/myproducts", server.listMyProduct)
		authRoutes.POST("/myproducts", server.createMyProduct)
		authRoutes.PUT("/myproduct/:product_id", server.updateMyProduct)

		authRoutes.POST("/myfundraise", server.createMyFundraise)
		authRoutes.GET("/myfundraise", server.getMyFundraise) 
		authRoutes.PUT("/exitmyfundraise", server.exitMyFundraise)
	
		authRoutes.POST("/transfer", server.createTransfer)
		server.router = router
}
// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}