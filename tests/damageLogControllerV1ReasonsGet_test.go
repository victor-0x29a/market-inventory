package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/market-inventory/constants"
	"github.com/stretchr/testify/assert"
)

func TestGetApiReasonsMapDamageLogControllerV1MustSuccess(t *testing.T) {
	app := Setup()

	req, _ := http.NewRequest("GET", "/v1/damage-log/reasons", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, res.StatusCode, 200)

	responseBody, _ := io.ReadAll(res.Body)

	var response map[int]string
	json.Unmarshal(responseBody, &response)

	assert.Equal(t, response[1], constants.ReasonsMap[1])
}
