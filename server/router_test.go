package server_test

import (
	"bytes"
	"crud-challenge/dto"
	"crud-challenge/server"
	"crud-challenge/storage"
	"encoding/json"
	"github.com/golobby/container/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	wagerList = []*storage.Wager{
		{
			Model: gorm.Model{
				ID: 1,
			},
			TotalWagerValue: 100,
			Odds: 2,
			SellingPercentage: 100,
			SellingPrice: 200,
		},
	}
)

func init() {
	container.Singleton(func() storage.IWagerDAO {
		mockStorage := &storage.MockIWagerDAO{}
		mockStorage.On("List", mock.Anything, 1, 10).
			Return(wagerList, nil)

		mockStorage.On("Create", mock.Anything, mock.Anything).
			Return(wagerList[0], nil)

		return mockStorage
	})
}

func TestListHappy(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/wagers?page=1&limit=10", nil)
	response := executeRequest(req)

	wagerListResp := []*dto.Wager{
		wagerList[0].ToDto(),
	}

	responseExpect, _ := json.Marshal(wagerListResp)

	assert.Equal(t, string(responseExpect), response.Body.String())

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCreateWagerHappy(t *testing.T) {
	var jsonStr = []byte(`{
	"total_wager_value": 100,
	"odds": 2,
	"selling_percentage": 100,
	"selling_price": 200
}`)

	req, _ := http.NewRequest("POST", "/v1/wagers", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)

	responseExpect, _ := json.Marshal(wagerList[0].ToDto(),)

	assert.Equal(t, string(responseExpect), response.Body.String())

	checkResponseCode(t, http.StatusCreated, response.Code)
}

func TestCreateWagerNotMonetary(t *testing.T) {
	var jsonStr = []byte(`{
	"total_wager_value": 100,
	"odds": 2,
	"selling_percentage": 100,
	"selling_price": 200.222
}`)

	req, _ := http.NewRequest("POST", "/v1/wagers", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)

	responseExpect, _ := json.Marshal(map[string]string{
		"message": "SellingPrice is not in monetary",
	})

	assert.Equal(t, string(responseExpect), response.Body.String())

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	router := server.InternalRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
