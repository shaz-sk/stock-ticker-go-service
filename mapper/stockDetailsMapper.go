package mapper

import (
	"fmt"
	"log"
	"maps"
	"math"
	"slices"
	"stock-ticker-go-service/data"
	"strconv"
)

// StockDetailsMapper is responsible for mapping information received from Alpha Vantage to API response.
//
// Fields:
// - ndays: The number of days for processing stock data.
type StockDetailsMapper struct {
	ndays int
}

func NewStockDetailsMapper(ndays int) *StockDetailsMapper {
	return &StockDetailsMapper{
		ndays: ndays,
	}
}

func (s *StockDetailsMapper) MapToStockDetails(stockData data.StockData) data.StockDetails {
	dailyClosingPrice := getDailyClosingPrice(stockData.TimeSeries, s.ndays, 30)
	return data.StockDetails{
		Information:         stockData.Information,
		Symbol:              stockData.MetaData.Symbol,
		DailyClosingPrice:   dailyClosingPrice,
		AverageClosingPrice: averagePrice(dailyClosingPrice, s.ndays),
		AveragePeriod:       s.ndays,
	}
}

func getDailyClosingPrice(dailyDataMap map[string]data.DailyData, ndays int, maxDays int) map[string]float64 {
	// Ensure ndays is not greater than maxDays
	if ndays > maxDays {
		ndays = maxDays
	}

	result := make(map[string]float64)

	// Iterate over the first ndays elements in the dailyDataMap
	count := 0
	sortedKeys := slices.Sorted(maps.Keys(dailyDataMap))
	for index := len(dailyDataMap) - 1; index >= 0; index-- {
		if count >= ndays {
			break
		}
		fmt.Println(sortedKeys[index])
		key := sortedKeys[index]
		floatValue, err := strconv.ParseFloat(dailyDataMap[key].Close, 64)
		if err != nil {
			log.Fatal(err)
		}
		result[key] = floatValue
		count++
	}
	return result
}

func averagePrice(dailyPrice map[string]float64, ndays int) float64 {
	var sum float64
	for _, price := range dailyPrice {
		sum += price
	}
	return math.Round(sum / float64(ndays))
}
