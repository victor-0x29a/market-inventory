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

func TestGetWithPaginationApiProductControllerV1MustSuccess(t *testing.T) {
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

	req, _ = http.NewRequest("GET", "/v1/product", nil)

	res, err = app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiPaginationResponse
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)
	assert.Equal(t, response.ItemsCount, int64(1))
	assert.Equal(t, response.TotalPages, int64(1))
	assert.NotEmpty(t, response.Records)

	raw := response.Records.([]interface{})

	firstRecord := raw[0].(map[string]interface{})

	assert.Equal(t, firstRecord["Title"], payload.Title)
}

func TestGetWithPaginationApiProductControllerV1MustSuccessWhenHavePaginatrionQueryParams(t *testing.T) {
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

	assert.Equal(t, res.StatusCode, 201)

	assert.Nil(t, err)

	// item 2
	payload = dtos.CreateProductDTO{
		Title:             "Cookie",
		Description:       nil,
		Price:             350,
		InventoryQuantity: 1,
	}

	body, _ = json.Marshal(payload)

	req, _ = http.NewRequest("POST", "/v1/product", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)
	// =====

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 201)

	req, _ = http.NewRequest("GET", "/v1/product?Page=1&PerPage=1", nil)

	res, err = app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiPaginationResponse
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)
	assert.Equal(t, response.ItemsCount, int64(2))
	assert.Equal(t, response.TotalPages, int64(2))
	assert.NotEmpty(t, response.Records)

	raw := response.Records.([]interface{})

	firstRecord := raw[0].(map[string]interface{})

	assert.Equal(t, firstRecord["Title"], "Fish")

	// second page

	req, _ = http.NewRequest("GET", "/v1/product?Page=2&PerPage=1", nil)

	res, err = app.Test(req)

	responseBody, _ = io.ReadAll(res.Body)

	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)
	assert.Equal(t, response.ItemsCount, int64(2))
	assert.Equal(t, response.TotalPages, int64(2))
	assert.NotEmpty(t, response.Records)

	raw = response.Records.([]interface{})

	firstRecord = raw[0].(map[string]interface{})

	assert.Equal(t, firstRecord["Title"], "Cookie")

	// unexists page

	req, _ = http.NewRequest("GET", "/v1/product?Page=3&PerPage=1", nil)

	res, err = app.Test(req)

	responseBody, _ = io.ReadAll(res.Body)

	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)
	assert.Equal(t, response.ItemsCount, int64(2))
	assert.Equal(t, response.TotalPages, int64(2))
	assert.Empty(t, response.Records)
}

func TestGetWithPaginationApiProductControllerV1MustSuccessWhenIsInvalidPaginationParams(t *testing.T) {
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

	req, _ = http.NewRequest("GET", "/v1/product?Page=kkkk", nil)

	res, err = app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiPaginationResponse
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)
	assert.Equal(t, response.ItemsCount, int64(1))
	assert.Equal(t, response.TotalPages, int64(1))
	assert.NotEmpty(t, response.Records)

	raw := response.Records.([]interface{})

	firstRecord := raw[0].(map[string]interface{})

	assert.Equal(t, firstRecord["Title"], payload.Title)

	// per page invalid

	req, _ = http.NewRequest("GET", "/v1/product?PerPage=kkkk", nil)

	res, err = app.Test(req)

	responseBody, _ = io.ReadAll(res.Body)

	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)
	assert.Equal(t, response.ItemsCount, int64(1))
	assert.Equal(t, response.TotalPages, int64(1))
	assert.NotEmpty(t, response.Records)

	raw = response.Records.([]interface{})

	firstRecord = raw[0].(map[string]interface{})

	assert.Equal(t, firstRecord["Title"], payload.Title)
}
