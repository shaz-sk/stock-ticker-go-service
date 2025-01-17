package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"stock-ticker-go-service/config"
	"stock-ticker-go-service/data"
	"testing"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) GetTimeSeriesData() (data.StockData, error) {
	args := m.Called()
	return args.Get(0).(data.StockData), args.Error(1)
}

func (m *MockClient) MapToStockDetails(stockData data.StockData) data.StockDetails {
	args := m.Called(stockData)
	return args.Get(0).(data.StockDetails)
}

func TestStockTickerService_GetClosingQuote(t *testing.T) {

	mockcfg := config.EnvConfig{
		Apikey: "test-api-key",
		Symbol: "GOOG",
		Ndays:  1,
		Url:    "https://api.example.com",
	}

	mockClient := new(MockClient)
	mockClient.On("GetTimeSeriesData").Return(getStockData(), nil)

	mockMapper := new(MockClient)
	mockMapper.On("MapToStockDetails", getStockData()).Return(expectedStockDetails(), nil)

	service := NewStockTickerService(mockcfg, mockClient, mockMapper)

	stockDetails := service.GetClosingQuote()

	assert.Equal(t, expectedStockDetails(), stockDetails)
	mockClient.AssertExpectations(t)
	mockMapper.AssertExpectations(t)
}

func getStockData() data.StockData {
	return data.StockData{
		MetaData: data.MetaData{
			Information:   "Daily Prices",
			Symbol:        "MSFT",
			LastRefreshed: "2024-11-22",
			OutputSize:    "Compact",
			TimeZone:      "US/Eastern",
		},
		TimeSeries: map[string]data.DailyData{
			"2024-11-22": {
				Open:   "d",
				High:   "411.365",
				Low:    "412.362",
				Close:  "414.361",
				Volume: "410.165",
			},
		},
	}
}

func expectedStockDetails() data.StockDetails {
	return data.StockDetails{
		Symbol:              "GOOG",
		AverageClosingPrice: 1500.0,
		AveragePeriod:       1,
		DailyClosingPrice: map[string]float64{
			"2024-11-22": 417.0,
			"2024-11-21": 412.87,
			"2024-11-20": 415.49,
		},
	}
}
