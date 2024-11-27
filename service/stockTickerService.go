package service

import (
	"log"
	"stock-ticker-go-service/config"
	"stock-ticker-go-service/data"
	"stock-ticker-go-service/mapper"
)

// StockTickerService handles the business logic for retrieving and processing stock ticker data.
//
// Fields:
// - logger: A logger instance used for logging informational messages, warnings,
//   and errors during the service's execution.
// - envconfig: Environment configuration required for AlphavantageClient and StockDetailsMapper.

type StockTickerService struct {
	logger             *log.Logger
	envconfig          *config.EnvConfig
	client             AlphaVantageClientInterface
	stockDetailsMapper mapper.StockDetailsMapperInterface
}

func NewStockTickerService(
	logger *log.Logger,
	config *config.EnvConfig,
	client AlphaVantageClientInterface,
	stockDetailsMapper mapper.StockDetailsMapperInterface,
) *StockTickerService {
	return &StockTickerService{
		logger:             logger,
		envconfig:          config,
		client:             client,
		stockDetailsMapper: stockDetailsMapper,
	}

}

func (s *StockTickerService) GetClosingQuote() data.StockDetails {
	timeSeriesData, _ := s.client.GetTimeSeriesData()
	return s.stockDetailsMapper.MapToStockDetails(timeSeriesData)
}
