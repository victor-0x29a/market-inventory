package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/constants"
	"github.com/stretchr/testify/assert"
)

func TestPostApiDamageLogControllerV1MustSuccess(t *testing.T) {
	app := Setup()

	// product creation

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

	// ----

	damageLogPayload := dtos.CreateDamageLogDTO{
		Reason:    1,
		Quantity:  1,
		ProductId: 1,
	}

	body, _ = json.Marshal(damageLogPayload)

	req, _ = http.NewRequest("POST", "/v1/damage-log", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 201)
}

func TestPostApiDamageLogControllerV1MustRejectWhenUnexistsProduct(t *testing.T) {
	app := Setup()

	damageLogPayload := dtos.CreateDamageLogDTO{
		Reason:    1,
		Quantity:  1,
		ProductId: 1,
	}

	body, _ := json.Marshal(damageLogPayload)

	req, _ := http.NewRequest("POST", "/v1/damage-log", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	assert.Nil(t, err)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 404)
	assert.Equal(t, response.Errors[0], constants.ErrProductNotFound.Error())
}

func TestPostApiDamageLogControllerV1MustRejectWhenIsInvalidReason(t *testing.T) {
	app := Setup()

	damageLogPayload := dtos.CreateDamageLogDTO{
		Reason:    9999,
		Quantity:  1,
		ProductId: 1,
	}

	body, _ := json.Marshal(damageLogPayload)

	req, _ := http.NewRequest("POST", "/v1/damage-log", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)

	assert.Nil(t, err)

	responseBody, _ := io.ReadAll(res.Body)

	var response dtos.ApiError
	json.Unmarshal(responseBody, &response)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 422)
	assert.Equal(t, response.Errors[0], "Key: 'CreateDamageLogDTO.Reason' Error:Field validation for 'Reason' failed on the 'is_valid_damage_reason' tag")

	// reason as 0

	damageLogPayload.Reason = 0

	body, _ = json.Marshal(damageLogPayload)

	req, _ = http.NewRequest("POST", "/v1/damage-log", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 422)
}
