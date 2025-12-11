package tests

import (
	"testing"

	dtos "github.com/market-inventory/DTOs"
	"github.com/market-inventory/utils"
)

func TestMustPass(t *testing.T) {
	var payload dtos.CreateProductDTO

	payload, statusCode := utils.ValidateStruct(&payload, make([]byte, 0))
}
