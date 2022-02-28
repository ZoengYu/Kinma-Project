package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	db "github.com/kinmaBackend/db/sqlc"
	"github.com/kinmaBackend/token"
	"github.com/kinmaBackend/util"
)

const (
	ErrInvalidOwner = "target request doesn't belong to the authenicated user"
)

var validCurrency validator.Func = func(filedLevel validator.FieldLevel) bool {
	if currency, ok := filedLevel.Field().Interface().(string); ok{
		return util.IsSupportedCurrency(currency)
	}

	return false
}

func validIsMyAccount(ctx *gin.Context, account db.Account, authPayload *token.Payload) bool{
	if account.Owner != authPayload.Username {
		err := errors.New(ErrInvalidOwner)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return false
	}
	return true
}