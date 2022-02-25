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


type transferRequest struct {
	FromAccountID int64 	`json:"from_account_id" binding:"required,min=1"`
	ToProductID 	int64 	`json:"to_product_id" binding:"required,min=1"`
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	//vaild if the account currency is match
	fromAccount, valid := server.validAccount(ctx, req.FromAccountID, req.Currency)
	if !valid{
		return
	}

	if fromAccount.Owner != authPayload.Username{
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	
	//check if fundraise was over
	targetFundraise, err := server.store.GetProductFundraise(ctx, req.ToProductID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if targetFundraise.Success{
		err := errors.New("fundraise project already over, thank you for the support")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.TransferParams{
		FromAccountID		: req.FromAccountID,
		ToProductID			: req.ToProductID,
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

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err :=server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}
		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	if account.Currency != currency{
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	return account, true

}