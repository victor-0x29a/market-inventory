package tests

import (
	"encoding/json"
	"log"
	"testing"

	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/utils"
	"github.com/stretchr/testify/assert"
)

func TestMustRejectValidateStructWhenNoHaveBody(t *testing.T) {
	var payload dtos.CreateProductDTO

	body := make([]byte, 0)

	data, statusCode := utils.ValidateStruct(&payload, body)

	assert.Equal(t, statusCode, 500)
	assert.Equal(t, data.Errors[0], "internal error")
}

func TestMustSuccesValidateStruct(t *testing.T) {
	var payload dtos.CreateProductDTO

	body, err := json.Marshal(dtos.CreateProductDTO{
		Title:             "top",
		Description:       nil,
		Price:             1050,
		InventoryQuantity: 5,
	})

	if err != nil {
		log.Fatal("error on creation of body")
	}

	data, statusCode := utils.ValidateStruct(&payload, body)

	assert.Equal(t, statusCode, 0)
	assert.Empty(t, data)
}

func TestMustRejectValidateStructAValidatorError(t *testing.T) {
	var payload dtos.CreateProductDTO

	body, err := json.Marshal(dtos.CreateProductDTO{
		Title:             "",
		Description:       nil,
		Price:             1050,
		InventoryQuantity: 5,
	})

	if err != nil {
		log.Fatal("error on creation of body")
	}

	data, statusCode := utils.ValidateStruct(&payload, body)

	assert.Equal(t, statusCode, 422)
	assert.Equal(t, data.Errors[0], "Key: 'CreateProductDTO.Title' Error:Field validation for 'Title' failed on the 'required' tag")
}
