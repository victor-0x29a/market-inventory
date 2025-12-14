package tests

import (
	"testing"

	"github.com/market-inventory/constants"
	"github.com/market-inventory/utils"
	"github.com/stretchr/testify/assert"
)

func TestMustSucessValidatePagination(t *testing.T) {
	apiPagination := utils.ValidatePagination("1", "10")

	assert.Equal(t, apiPagination.Page, 1)
	assert.Equal(t, apiPagination.PerPage, 10)
}

func TestValidatePaginationMustAcceptWhenIsNaN(t *testing.T) {
	apiPagination := utils.ValidatePagination("1a", "10s")

	assert.Equal(t, apiPagination.Page, constants.DefaultPage)
	assert.Equal(t, apiPagination.PerPage, constants.DefaultPerPage)
}
