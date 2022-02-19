package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kinmaBackend/db/sqlc"
)

type CreateFundraiseParams struct {
	ProductID      int64 `json:"product_id" binding:"required"`
	TargetAmount   int64 `json:"target_amount" binding:"required"`
	ProgressAmount int64 `json:"progress_amount"`
}

// Server expose method for API
func (server *Server) createFundraise(ctx *gin.Context){
	var req CreateFundraiseParams
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFundraiseParams{
		ProductID: req.ProductID,
		TargetAmount: req.TargetAmount,
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

type getFundraiseRequest struct {
	ProductID  int64  `uri:"id" binding:"required,min=1"`
}

// Server expose method for API
func (server *Server) getFundraise(ctx *gin.Context){
	var req getFundraiseRequest
	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//Implement the DB CRUD
	fundraise, err := server.store.GetFundraise(ctx, req.ProductID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, fundraise)
}

type exitFundraiseRequest struct {
	ProductID int64 `json:"product_id" binding:"required,min=1"`
	Success   *bool `json:"success" binding:"required"`
}

func (server *Server) exitFundraise(ctx *gin.Context){
	var req exitFundraiseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
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