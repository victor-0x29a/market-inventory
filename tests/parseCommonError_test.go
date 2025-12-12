package tests

import (
	"errors"
	"testing"

	"github.com/market-inventory/constants"
	"github.com/market-inventory/utils"
	"github.com/stretchr/testify/assert"
)

func TestMustSuccessParseCommonError(t *testing.T) {
	data, statusCode := utils.ParseCommonError(errors.New("unknown error"))

	assert.Equal(t, statusCode, 500)
	assert.Equal(t, data.Errors[0], constants.ErrInternal.Error())
}

func TestMustRejectWhenIsErrNotFoundProduct(t *testing.T) {
	data, statusCode := utils.ParseCommonError(constants.ErrProductNotFound)

	assert.Equal(t, statusCode, 404)
	assert.Equal(t, data.Errors[0], constants.ErrProductNotFound.Error())
}
