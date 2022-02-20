package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kinmaBackend/db/sqlc"
)


type transferRequest struct {
	FromAccountID int64 	`json:"from_account_id" binding:"required,min=1"`
	ToFundraiseID int64 	`json:"to_fundraise_id" binding:"required,min=1"`
	Amount 				int64 	`json:"amount" binding:"required,gt=0"`
	Currency			string	`json:"currency" binding:"required,currency"`
}

// Server expose method for API
func (server *Server) createTransfer(ctx *gin.Context){
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//vaild if the account currency is match
	if !server.validAccount(ctx, req.FromAccountID, req.Currency){
		return
	}

	arg := db.TransferParams{
		FromAccountID		: req.FromAccountID,
		ToFundraiseID		: req.ToFundraiseID,
		Amount					: req.Amount,
	}
	//Implement the DB CRUD
	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool{
	account, err :=server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency{
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true

}