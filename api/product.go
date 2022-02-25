package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kinmaBackend/db/sqlc"
	"github.com/kinmaBackend/token"
)


type createProducRequest struct {
	AccountID  			int64  		`json:"account_id" binding:"required"`
	Title      			string 		`json:"title" binding:"required"`
	Content    		 	string 		`json:"content" binding:"required"`
	ProductTagList 	[]string 	`json:"product_tag" binding:"required"`
}

// Server expose method for API
func (server *Server) createProduct (ctx *gin.Context){
	var req createProducRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := server.store.GetAccount(ctx, req.AccountID)
	if err != nil{
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Username != account.Owner {
		err := errors.New("account doesn't belong to the authenicated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.CreateProductParams{
		AccountID: req.AccountID,
		Title: req.Title,
		Content: req.Content,
		ProductTag: req.ProductTagList,
	}
	//Implement the DB CRUD
	product, err := server.store.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, product)
}

type getProductRequest struct {
	ProductID  int64  `uri:"id" binding:"required,min=1"`
}

// Server expose method for API
func (server *Server) getProduct(ctx *gin.Context){
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//Implement the DB CRUD
	product, err := server.store.GetAccountProduct(ctx, req.ProductID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}


	ctx.JSON(http.StatusOK, product)
}

type listMyProductRequest struct {
	AccountID int64 `form:"account_id" binding:"required,min=1"`
	PageID 		int32 `form:"page_id" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listMyProduct(ctx *gin.Context){
	var req listMyProductRequest
	if err := ctx.ShouldBindQuery(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	account, err := server.store.GetAccount(ctx, req.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	if authPayload.Username != account.Owner{
		err := errors.New("product doesn't belong to the authenicated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	
	arg := db.ListMyProductParams{
		AccountID	: req.AccountID,
		Limit			: req.PageSize,
		Offset		: (req.PageID -1) * req.PageSize,
	}
	products, err := server.store.ListMyProduct(ctx, arg)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}

type updateProductIDRequest struct {
	ProductID 	int64		`uri:"product_id" binding:"required,min=1"`
}


type updateProductRequest struct {
	AccountID				int64			`json:"account_id" binding:"required"`
	Title      			string 		`json:"title" binding:"required"`
	Content    			string 		`json:"content" binding:"required"`
	ProductTagList 	[]string 	`json:"product_tag" binding:"required"`
}

func (server *Server) updateProduct(ctx *gin.Context){
	var reqID updateProductIDRequest
	var req updateProductRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	account, err := server.store.GetAccount(ctx, req.AccountID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Username != account.Owner{
		err := errors.New("account doesn't belong to the authenicated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.UpdateProductDetailParams{
		ID: 				reqID.ProductID,
		Title: 			req.Title,
		Content: 		req.Content,
		ProductTag: req.ProductTagList,
	}

	updatedProduct, err := server.store.UpdateProductDetail(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}