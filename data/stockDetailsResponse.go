package data

type StockDetails struct {
	Information         string             `json:"information,omitempty"`
	Symbol              string             `json:"symbol,omitempty"`
	AverageClosingPrice float64            `json:"averageClosingPrice,omitempty"`
	AveragePeriod       int                `json:"averagePeriod,omitempty"`
	DailyClosingPrice   map[string]float64 `json:"dailyClosingPrice,omitempty"`
}
