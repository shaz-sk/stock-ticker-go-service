package config

import (
	"fmt"
	"os"
	"strconv"
)

// EnvConfig struct holds all the configuration settings for environment variables
type EnvConfig struct {
	Apikey string
	Symbol string
	Ndays  int
	Url    string
}

func NewConfig() (*EnvConfig, error) {
	apikey, present := os.LookupEnv("APIKEY")
	if !present {
		fmt.Println("APIKEY environment variable is not set")
		apikey = "test"
	}

	symbol, present := os.LookupEnv("SYMBOL")
	if !present {
		fmt.Println("SYMBOL environment variable is not set, using default value: MSFT")
		symbol = "MSFT"
	}

	ndays, err := strconv.Atoi(os.Getenv("NDAYS"))
	if err != nil {
		fmt.Println("NDAYS environment variable is not set, using default value: 3")
		ndays = 3
	}

	url := "https://www.alphavantage.co/query"

	return &EnvConfig{
		Apikey: apikey,
		Symbol: symbol,
		Ndays:  ndays,
		Url:    url,
	}, nil
}
