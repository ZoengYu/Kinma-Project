package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/kinmaBackend/util"
)

var validCurrency validator.Func = func(filedLevel validator.FieldLevel) bool {
	if currency, ok := filedLevel.Field().Interface().(string); ok{
		return util.IsSupportedCurrency(currency)
	}

	return false
}