package main

import (
	"log"
	"net/http"
	"os"
	"stock-ticker-go-service/config"
	"stock-ticker-go-service/handler"
	"stock-ticker-go-service/mapper"
	"stock-ticker-go-service/service"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "stock-ticker", log.LstdFlags)
	cfg, _ := config.NewConfig()
	client := service.NewApiVantageClient(cfg)
	stockMapper := mapper.NewStockDetailsMapper(cfg.Ndays)
	stockService := service.NewStockTickerService(logger, cfg, client, stockMapper)

	stockTickerHandler := handler.NewStockTickerHandler(logger, stockService)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/stock/report", stockTickerHandler)
	server := &http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	logger.Fatal(server.ListenAndServe())
}
