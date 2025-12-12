package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiIndexRouteMustSuccess(t *testing.T) {
	app := Setup()

	req, _ := http.NewRequest("GET", "/", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 200)
}
