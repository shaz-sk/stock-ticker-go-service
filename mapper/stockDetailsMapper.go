package mapper

import (
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

type StockDetailsMapperInterface interface {
	MapToStockDetails(stockData data.StockData) data.StockDetails
}
type StockDetailsMapper struct {
	ndays int
}

func NewStockDetailsMapper(ndays int) *StockDetailsMapper {
	return &StockDetailsMapper{
		ndays: ndays,
	}
}

func (s *StockDetailsMapper) MapToStockDetails(stockData data.StockData) data.StockDetails {
	filteredDailyData := getNDaysDailyData(stockData.TimeSeries, s.ndays, 30)
	dailyClosingPrice := getDailyClosingPrice(filteredDailyData)
	return data.StockDetails{
		Information:         stockData.Information,
		Symbol:              stockData.MetaData.Symbol,
		DailyClosingPrice:   dailyClosingPrice,
		AverageClosingPrice: averagePrice(dailyClosingPrice, s.ndays),
		AveragePeriod:       s.ndays,
	}
}

func getNDaysDailyData(dailyDataMap map[string]data.DailyData, ndays int, maxDays int) map[string]data.DailyData {
	// Ensure ndays is not greater than maxDays
	if ndays > maxDays {
		ndays = maxDays
	}

	result := make(map[string]data.DailyData)

	// Iterate over the first ndays elements in the dailyDataMap
	count := 0
	sortedKeys := slices.Sorted(maps.Keys(dailyDataMap))
	for index := len(dailyDataMap) - 1; index >= 0; index-- {
		if count >= ndays {
			break
		}
		key := sortedKeys[index]
		result[key] = dailyDataMap[key]
		count++
	}
	return result
}

func getDailyClosingPrice(dailyDataMap map[string]data.DailyData) map[string]float64 {
	result := make(map[string]float64)
	for key, value := range dailyDataMap {
		floatValue, err := strconv.ParseFloat(value.Close, 64)
		if err != nil {
			log.Printf("error parsing DailyClosingPrice data for %v%v", key, err)
		}
		result[key] = floatValue
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
