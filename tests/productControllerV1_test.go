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

	assert.Equal(t, res.StatusCode, 201)
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

func TestUpdateApiProductControllerV1MustSuccess(t *testing.T) {
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

	payload.Title = "Fish 2"

	body, _ = json.Marshal(payload)

	req, _ = http.NewRequest("PATCH", "/v1/product/1", bytes.NewReader(body))

	res, err = app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 204)

	req, _ = http.NewRequest("GET", "/v1/product/1", nil)

	res, err = app.Test(req)

	assert.Nil(t, err)

	responseBody, _ := io.ReadAll(res.Body)

	var response database.Product
	json.Unmarshal(responseBody, &response)

	assert.Equal(t, res.StatusCode, 200)
	assert.Equal(t, response.Title, "Fish 2")
	assert.NotEqual(t, response.Title, "Fish")
	assert.Equal(t, response.Price, payload.Price)
	assert.Equal(t, response.InventoryQuantity, payload.InventoryQuantity)
}

func TestUpdateApiProductControllerV1MustRejectWhenInvalidProductId(t *testing.T) {
	app := Setup()

	payload := dtos.CreateProductDTO{
		Title:             "Fish",
		Description:       nil,
		Price:             1050,
		InventoryQuantity: 1,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PATCH", "/v1/product/1k", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 422)
	assert.Equal(t, response.Errors[0], "Key: 'FetchProductDTO.ID' Error:Field validation for 'ID' failed on the 'required' tag")
}

func TestUpdateApiProductControllerV1MustRejectWhenInvalidDataProp(t *testing.T) {
	app := Setup()

	payload := dtos.CreateProductDTO{
		Title:             "Fish",
		Description:       nil,
		Price:             -1050,
		InventoryQuantity: 1,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PATCH", "/v1/product/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 422)
	assert.Equal(t, response.Errors[0], "Key: 'UpdateProductDTO.Price' Error:Field validation for 'Price' failed on the 'gt' tag")
}

func TestUpdateApiProductControllerV1MustRejectWhenUnexistsProduct(t *testing.T) {
	app := Setup()

	payload := dtos.CreateProductDTO{
		Title:             "Fish",
		Description:       nil,
		Price:             1050,
		InventoryQuantity: 1,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PATCH", "/v1/product/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 404)
	assert.Equal(t, response.Errors[0], "product not found")
}
