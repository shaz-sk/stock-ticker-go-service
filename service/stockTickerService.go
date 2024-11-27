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
	logger    *log.Logger
	envconfig *config.EnvConfig
}

func NewStockTickerService(
	logger *log.Logger,
	config *config.EnvConfig,
) *StockTickerService {
	return &StockTickerService{
		logger:    logger,
		envconfig: config,
	}

}

func (s *StockTickerService) GetClosingQuote() data.StockDetails {
	timeSeriesData := NewApiVantageClient(s.envconfig).GetTimeSeriesData()
	stockDetailsMapper := mapper.NewStockDetailsMapper(s.envconfig.Ndays)
	return stockDetailsMapper.MapToStockDetails(timeSeriesData)
}
