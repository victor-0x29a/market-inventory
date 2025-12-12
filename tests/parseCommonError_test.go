package tests

import (
	"errors"
	"testing"

	"github.com/market-inventory/utils"
	"github.com/stretchr/testify/assert"
)

func TestMustSuccessParseCommonError(t *testing.T) {
	data, statusCode := utils.ParseCommonError(errors.New("market error"))

	assert.Equal(t, statusCode, 500)
	assert.Equal(t, data.Errors[0], "internal error")
}
