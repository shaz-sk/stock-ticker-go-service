package service

import (
	"stock-ticker-go-service/config"
	"stock-ticker-go-service/data"
	"stock-ticker-go-service/mapper"
)

// StockTickerService handles the business logic for retrieving and processing stock ticker data.
//
// Fields:
//   and errors during the service's execution.
// - envconfig: Environment configuration required for AlphavantageClient and StockDetailsMapper.

type StockTickerService struct {
	envconfig          config.EnvConfig
	client             AlphaVantageClientInterface
	stockDetailsMapper mapper.StockDetailsMapperInterface
}

func NewStockTickerService(
	config config.EnvConfig,
	client AlphaVantageClientInterface,
	stockDetailsMapper mapper.StockDetailsMapperInterface,
) *StockTickerService {
	return &StockTickerService{
		envconfig:          config,
		client:             client,
		stockDetailsMapper: stockDetailsMapper,
	}

}

func (s *StockTickerService) GetClosingQuote() data.StockDetails {
	timeSeriesData, _ := s.client.GetTimeSeriesData()
	return s.stockDetailsMapper.MapToStockDetails(timeSeriesData)
}
