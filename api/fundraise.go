package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kinmaBackend/db/sqlc"
	"github.com/kinmaBackend/token"
)

type createFundraiseParams struct {
	ProductID  		 int64 `json:"product_id" binding:"required,min=1"`
	TargetAmount   int64 `json:"target_amount" binding:"required"`
	ProgressAmount int64 `json:"progress_amount"`
}

// Server expose method for API
func (server *Server) createMyFundraise(ctx *gin.Context){
	var req createFundraiseParams

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//check if fundraise already exist
	_, err := server.store.GetProductFundraise(ctx, req.ProductID)
	if err != sql.ErrNoRows{
		err := errors.New("this product already have fundraise")
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return 
	}

	product, err := server.store.GetProduct(ctx, req.ProductID)
	if err != nil{
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, product.AccountID)
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

	arg := db.CreateFundraiseParams{
		ProductID: 			req.ProductID,
		TargetAmount: 	req.TargetAmount,
		ProgressAmount: req.ProgressAmount,
	}
	//Implement the DB CRUD
	fundraise, err := server.store.CreateFundraise(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, fundraise)
}

type getMyFundraiseRequest struct {
	ProductID 	int64 	`json:"product_id" binding:"required"`
}

// Server expose method for API
func (server *Server) getMyFundraise(ctx *gin.Context){
	var req getMyFundraiseRequest

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//check if productID belong the owner
	product, err := server.store.GetProduct(ctx, req.ProductID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, product.AccountID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	//valid if the request target is belong to the login user
	if !validIsMyAccount(ctx, account, authPayload){
		return
}

	//Implement the DB CRUD
	fundraise, err := server.store.GetProductFundraise(ctx, req.ProductID)
	if err != nil {
		if err == sql.ErrNoRows{
			err := errors.New("cannot find the fundraise of the product")
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
	
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, fundraise)
}

type exitFundraiseRequest struct {
	ProductID  int64  `json:"product_id" binding:"required,min=1"`
	// pointer here to allow user type false as input 
	Success   		*bool `json:"success"`
}

func (server *Server) exitMyFundraise(ctx *gin.Context){
	var req exitFundraiseRequest

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.GetProduct(ctx, req.ProductID)
	if err != nil{
		if err == sql.ErrNoRows{
			err := errors.New("cannot find the product")
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//get product account owner
	account, err := server.store.GetAccount(ctx, product.AccountID)
	if err != nil{
		if err == sql.ErrNoRows{
			err := errors.New("cannot find the account")
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
	//check if fundraise exist
	fundraise, err := server.store.GetProductFundraise(ctx, req.ProductID)
	if err != nil{
		if err == sql.ErrNoRows{
			err := errors.New("cannot find the fundraise of product")
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	//init success as false
	success := new(bool)
	fmt.Println(*success,fundraise.TargetAmount, fundraise.ProgressAmount)
	//pointer has to point to something, assign value to another address
	if (fundraise.TargetAmount < fundraise.ProgressAmount){
		//assign value to the success pointer
		*success = true
		req.Success = success
		fmt.Println("pass")
	} else {
		req.Success = success
		fmt.Println("not pass")
	}

	arg := db.ExitFundraiseParams{
		ProductID	: req.ProductID,
		Success		: *req.Success,
	}

	exitFundraise, err := server.store.ExitFundraise(ctx, arg)
	if err != nil{
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, exitFundraise)
}

type addFundraiseProgressAmountRequest struct {
	// 0 is considered a unvalid value, pointer here will allow user to input 0
	Amount 			*int64 	`json:"amount" binding:"required"`
	ProductID   int64 	`json:"product_id" binding:"required"`
}

func (server *Server) addFundraise(ctx *gin.Context){
	var req addFundraiseProgressAmountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddFundraiseProgressAmountParams{
		Amount: *req.Amount,
		ID:			req.ProductID,
	}

	if (arg.Amount == 0){
		err := errors.New("cannot add invalid amount: 0")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	fundraise, err := server.store.AddFundraiseProgressAmount(ctx, arg)
	if err != nil{
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, fundraise)
}