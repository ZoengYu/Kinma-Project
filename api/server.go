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

		authRoutes.POST("/accounts", server.createAccount)
		authRoutes.GET("/accounts/:id", server.getAccount)
		authRoutes.GET("/accounts", server.listAccount)

		authRoutes.GET("/myproducts/:id", server.getProduct)
		authRoutes.GET("/myproducts", server.listMyProduct)
		authRoutes.POST("/myproducts", server.createProduct)
		authRoutes.PUT("/myproduct/:product_id", server.updateProduct)

		authRoutes.GET("/myproduct/:product_id/fundraise", server.getFundraise)
		authRoutes.POST("/myproduct/:product_id/fundraise", server.createFundraise)
		authRoutes.PUT("myproduct/:product_id/exitfundraise", server.exitFundraise)
	
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