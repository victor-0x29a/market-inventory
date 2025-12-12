package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	dtos "github.com/market-inventory/DTOs"
	"github.com/stretchr/testify/assert"
)

func TestPostApiProductControllerV1MustSuccess(t *testing.T) {
	app := Setup()

	payload := dtos.CreateProductDTO{
		Title:             "Fish",
		Description:       nil,
		Price:             350,
		InventoryQuantity: 1,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/v1/product", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 204)
}

func TestPostApiProductControllerV1MustRejectWhenHaveFieldMissing(t *testing.T) {
	app := Setup()

	payload := dtos.CreateProductDTO{
		Title:             "",
		Description:       nil,
		Price:             350,
		InventoryQuantity: 1,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/v1/product", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 422)
	assert.Equal(t, response.Errors[0], "Key: 'CreateProductDTO.Title' Error:Field validation for 'Title' failed on the 'required' tag")
}

func TestPostApiProductControllerV1MustRejectWhenPriceLessThanZero(t *testing.T) {
	app := Setup()

	payload := dtos.CreateProductDTO{
		Title:             "Fish",
		Description:       nil,
		Price:             -1050,
		InventoryQuantity: 1,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/v1/product", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 422)
	assert.Equal(t, response.Errors[0], "Key: 'CreateProductDTO.Price' Error:Field validation for 'Price' failed on the 'gt' tag")
}

func TestPostApiProductControllerV1MustRejectWhenQuantityLessThanZero(t *testing.T) {
	app := Setup()

	payload := dtos.CreateProductDTO{
		Title:             "Fish",
		Description:       nil,
		Price:             1050,
		InventoryQuantity: -1050,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/v1/product", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 422)
	assert.Equal(t, response.Errors[0], "Key: 'CreateProductDTO.InventoryQuantity' Error:Field validation for 'InventoryQuantity' failed on the 'gt' tag")
}
