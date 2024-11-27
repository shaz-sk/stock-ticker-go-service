package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"stock-ticker-go-service/service"
)

// StockTickerHandler is responsible for handling HTTP requests related to stock ticker information.
// Fields:
//   - stockTickerService: StockTickerService that provides the business logic for retrieving and processing stock ticker data.
type StockTickerHandler struct {
	stockTickerService *service.StockTickerService
}

func NewStockTickerHandler(tickerService *service.StockTickerService) *StockTickerHandler {
	return &StockTickerHandler{tickerService}

}

func (s *StockTickerHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	response := s.stockTickerService.GetClosingQuote()

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(writer, "Unable to process the request", http.StatusInternalServerError)
	}
}
