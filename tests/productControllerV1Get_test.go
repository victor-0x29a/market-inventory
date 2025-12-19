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

func TestGetApiProductControllerV1MustSuccess(t *testing.T) {
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
	assert.Equal(t, response.Title, payload.Title)
	assert.Equal(t, response.Price, payload.Price)
	assert.Equal(t, response.InventoryQuantity, payload.InventoryQuantity)
}

func TestGetApiProductControllerV1MustRejectWhenUnexists(t *testing.T) {
	app := Setup()

	req, _ := http.NewRequest("GET", "/v1/product/999", nil)

	res, err := app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 404)
	assert.Equal(t, response.Errors[0], constants.ErrProductNotFound.Error())
}

func TestGetApiProductControllerV1MustRejectWhenIdIsNaN(t *testing.T) {
	app := Setup()

	req, _ := http.NewRequest("GET", "/v1/product/1a", nil)

	res, err := app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 422)
	assert.Equal(t, response.Errors[0], "Key: 'FetchProductDTO.ID' Error:Field validation for 'ID' failed on the 'required' tag")
}
