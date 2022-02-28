package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kinmaBackend/db/sqlc"
	"github.com/kinmaBackend/token"
	"github.com/lib/pq"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

// Server expose method for API
func (server *Server) createAccount(ctx *gin.Context){
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//get authPayload from gin.context
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAccountParams{
		Owner		: authPayload.Username,
		Currency: req.Currency,
	}
	//Implement the DB CRUD
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return 
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID 			int64		`uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context){
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
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
	if !validIsMyAccount(ctx, account, authPayload){
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID 		int32	`form:"page_id" binding:"required,min=1"`
	PageSize 	int32	`form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listMyAccount(ctx *gin.Context){
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListAccountParams{
		Owner	: authPayload.Username,
		Limit	: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := server.store.ListAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}