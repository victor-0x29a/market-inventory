package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustSuccessParseCommonError(t *testing.T) {
	data, statusCode := ParseCommonError(errors.New("market error"))

	assert.Equal(t, statusCode, 500)
	assert.Equal(t, data.Errors[0], "internal error")
}
