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

type createFundraiseProductIDParams struct {
	ProductID  int64  `uri:"product_id" binding:"required,min=1"`
}

type createFundraiseParams struct {
	AccountID			 int64 `json:"account_id" binding:"required"`
	TargetAmount   int64 `json:"target_amount" binding:"required"`
	ProgressAmount int64 `json:"progress_amount"`
}

// Server expose method for API
func (server *Server) createFundraise(ctx *gin.Context){
	var req createFundraiseParams
	var reqProductID createFundraiseProductIDParams

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqProductID); err != nil{
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

	arg := db.CreateFundraiseParams{
		ProductID: 			reqProductID.ProductID,
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

type getFundraiseProductRequest struct {
	ProductID  int64  `uri:"product_id" binding:"required,min=1"`
}

// Server expose method for API
func (server *Server) getFundraise(ctx *gin.Context){
	var req getFundraiseProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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

type exitFundraiseProductIDRequest struct {
	ProductID  int64  `uri:"product_id" binding:"required,min=1"`
}

type exitFundraiseRequest struct {
	AccountID 		int64	`json:"account_id" binding:"required"`
	// pointer here to allow user type false as input 
	Success   		*bool `json:"success"`
}

func (server *Server) exitFundraise(ctx *gin.Context){
	var req exitFundraiseRequest
	var reqProductID exitFundraiseProductIDRequest

	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&req); err != nil{
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
	//check if fundraise exist
	fundraise, err := server.store.GetProductFundraise(ctx, reqProductID.ProductID)
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
		ProductID	: reqProductID.ProductID,
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