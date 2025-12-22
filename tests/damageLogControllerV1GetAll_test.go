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

func TestGetWithPaginationApiDamageLogControllerV1MustSuccess(t *testing.T) {
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
		Quantity:  1,
		Reason:    1,
		ProductId: 1,
	}

	body, _ = json.Marshal(damageLogPayload)

	req, _ = http.NewRequest("POST", "/v1/damage-log", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err = app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 201)

	req, _ = http.NewRequest("GET", "/v1/damage-log", nil)

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

	assert.Equal(t, firstRecord["Quantity"], float64(damageLogPayload.Quantity))
}
