package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"stock-ticker-go-service/config"
	"stock-ticker-go-service/data"
)

// AlphaVantageClient is responsible for calling Alpha Vantage API to retrieve stock data.
//
// Fields:
// - envconfig: Environment configuration required to connect to Alpha Vantage API.

type AlphaVantageClientInterface interface {
	GetTimeSeriesData() (data.StockData, error)
}

type AlphaVantageClient struct {
	envconfig *config.EnvConfig
}

func NewApiVantageClient(envconfig *config.EnvConfig) *AlphaVantageClient {
	return &AlphaVantageClient{envconfig: envconfig}
}

func (a *AlphaVantageClient) GetTimeSeriesData() (data.StockData, error) {

	var apiResponse data.StockData
	// Make the GET request
	fmt.Printf("Retrieving data from Alphavantage with symbol %v\n", a.envconfig.Symbol)
	resp, err := http.Get(getUrl(a.envconfig))
	if err != nil {
		fmt.Errorf("alphavantage service unavailable %v", err)
		return apiResponse, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("alphavantage service unavailable %v", err)
		return apiResponse, err
	}

	// Map the JSON response to the struct
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Errorf("alphavantage response parse error %v", err)
	}
	//fmt.Println("Response: ", string(body))
	return apiResponse, err

}

func getUrl(envconfig *config.EnvConfig) string {
	baseURL := "https://www.alphavantage.co/query"

	params := url.Values{}
	params.Add("symbol", envconfig.Symbol)
	params.Add("apikey", envconfig.Apikey)
	params.Add("function", "TIME_SERIES_DAILY")

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}
