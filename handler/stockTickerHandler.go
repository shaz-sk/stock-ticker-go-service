package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"stock-ticker-go-service/service"
)

// StockTickerHandler is responsible for handling HTTP requests related to stock ticker information.
// Fields:
//   - logger: A logger instance used to log events, errors, or other relevant information.
//   - stockTickerService: StockTickerService that provides the business logic for retrieving and processing stock ticker data.
type StockTickerHandler struct {
	logger             *log.Logger
	stockTickerService *service.StockTickerService
}

func NewStockTickerHandler(logger *log.Logger, tickerService *service.StockTickerService) *StockTickerHandler {
	return &StockTickerHandler{logger, tickerService}

}

func (s *StockTickerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	response := s.stockTickerService.GetClosingQuote()

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		s.logger.Println("Error encoding JSON:", err)
		http.Error(writer, "Unable to process the request", http.StatusInternalServerError)
	}
}
