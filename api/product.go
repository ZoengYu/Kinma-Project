package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kinmaBackend/db/sqlc"
	"github.com/kinmaBackend/token"
)


type createMyProducRequest struct {
	AccountID  			int64  		`json:"account_id" binding:"required"`
	Title      			string 		`json:"title" binding:"required"`
	Content    		 	string 		`json:"content" binding:"required"`
	ProductTagList 	[]string 	`json:"product_tag" binding:"required"`
}

// Server expose method for API
func (server *Server) createMyProduct (ctx *gin.Context){
	var req createMyProducRequest
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
	
	//valid if the request is belong to the login user
	if !validIsMyAccount(ctx, account, authPayload){
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

type getMyProductIDRequest struct {
	ProductID  int64  `uri:"product_id" binding:"required,min=1"`
}

type getMyAccountProductRequest struct {
	AccountID  int64  `json:"account_id" binding:"required,min=1"`
}

// Server expose method for API
func (server *Server) getMyProduct(ctx *gin.Context){
	var reqProductID getMyProductIDRequest
	var reqOwner getMyAccountProductRequest

	if err := ctx.ShouldBindJSON(&reqOwner); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqProductID); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	product, err := server.store.GetProduct(ctx, reqProductID.ProductID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	productOwner, err := server.store.GetAccount(ctx, product.AccountID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	//valid if the request product is belong to the login user
	if !validIsMyAccount(ctx, productOwner, authPayload){
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
	
	//valid if the request is belong to the login user
	if !validIsMyAccount(ctx, account, authPayload){
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

type updateMyProductIDRequest struct {
	ProductID 	int64		`uri:"product_id" binding:"required,min=1"`
}


type updateMyProductRequest struct {
	AccountID				int64			`json:"account_id" binding:"required"`
	Title      			string 		`json:"title" binding:"required"`
	Content    			string 		`json:"content" binding:"required"`
	ProductTagList 	[]string 	`json:"product_tag" binding:"required"`
}

func (server *Server) updateMyProduct(ctx *gin.Context){
	var reqProduct updateMyProductIDRequest
	var req updateMyProductRequest
	if err := ctx.ShouldBindUri(&reqProduct); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	product, err := server.store.GetProduct(ctx, reqProduct.ProductID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	productOwner, err := server.store.GetAccount(ctx, product.AccountID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	//valid if the request is belong to the login user
	if !validIsMyAccount(ctx, productOwner, authPayload){
		return
}

	arg := db.UpdateProductDetailParams{
		ID: 				reqProduct.ProductID,
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