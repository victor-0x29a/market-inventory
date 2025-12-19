package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/constants"
	"github.com/market-inventory/database"
	"github.com/stretchr/testify/assert"
)

func TestDeleteApiProductControllerV1MustSuccess(t *testing.T) {
	app := Setup()

	payload := dtos.CreateProductDTO{
		Title:             "Fish",
		Description:       nil,
		Price:             1050,
		InventoryQuantity: 1,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/v1/product", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 201)

	req, _ = http.NewRequest("GET", "/v1/product/1", nil)

	res, err = app.Test(req)

	assert.Nil(t, err)

	responseBody, _ := io.ReadAll(res.Body)

	var response database.Product
	json.Unmarshal(responseBody, &response)

	assert.Equal(t, res.StatusCode, 200)

	// delete

	req, _ = http.NewRequest("DELETE", "/v1/product/1", nil)

	res, err = app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 204)

	// --

	req, _ = http.NewRequest("GET", "/v1/product/1", nil)

	res, err = app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 404)

	responseBody, _ = io.ReadAll(res.Body)

	var responseErr dtos.ApiError

	json.Unmarshal(responseBody, &responseErr)

	assert.Nil(t, err)

	assert.Equal(t, responseErr.Errors[0], constants.ErrProductNotFound.Error())
}

func TestDeleteApiProductControllerV1MustRejectWhenUnexists(t *testing.T) {
	app := Setup()

	req, _ := http.NewRequest("DELETE", "/v1/product/1", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 404)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, response.Errors[0], constants.ErrProductNotFound.Error())
}

func TestDeleteApiProductControllerV1MustRejectWhenIdIsInvalid(t *testing.T) {
	app := Setup()

	req, _ := http.NewRequest("DELETE", "/v1/product/1a", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 422)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, response.Errors[0], "Key: 'FetchProductDTO.ID' Error:Field validation for 'ID' failed on the 'required' tag")
}
