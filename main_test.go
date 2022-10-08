package main

import (
	"encoding/json"
	"github.com/FloatKasemtan/endpoints/covid"
	"github.com/FloatKasemtan/types"
	"github.com/FloatKasemtan/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestParseJsonFunction(t *testing.T) {
	// read mock data
	jsonFile, err := os.Open("./test_data/raw.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	// convert file to byte array
	mockResponse, _ := ioutil.ReadAll(jsonFile)
	// initialize expect result
	expect := new(types.CovidResponse)
	json.Unmarshal([]byte(mockResponse), &expect)

	dataUrl := "http://static.wongnai.com/devinterview/covid-cases.json"

	// initialize actual result
	actual := new(types.CovidResponse)
	if err := utils.GetJson[*types.CovidResponse](dataUrl, actual); err != nil {
		panic(err)
	}

	assert.Equal(t, *expect, *actual)
}

func TestHomepageHandler(t *testing.T) {
	const path = "/covid/summary"

	// read mock response
	jsonFile, err := os.Open("./test_data/summary.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	// convert file to byte array
	mockResponse, _ := ioutil.ReadAll(jsonFile)
	// initialize expect result
	var expect map[string]interface{}
	json.Unmarshal([]byte(mockResponse), &expect)

	// set up router
	r := SetUpRouter()
	r.GET(path, covid.SummaryHandler)
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// parse body to byte array
	responseData, _ := ioutil.ReadAll(w.Body)
	// initialize actual result
	var actual map[string]interface{}
	json.Unmarshal([]byte(responseData), &actual)

	assert.Equal(t, expect, actual)
	assert.Equal(t, http.StatusOK, w.Code)
}
