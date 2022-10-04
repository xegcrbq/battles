package binance

import (
	"strconv"
	"strings"
)

type binanceAnswer struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	Price    float64
}

func filterAnswers(ba []binanceAnswer, ticker string) []binanceAnswer {
	result := make([]binanceAnswer, len(ba)/2)
	for i := range ba {
		if strings.Contains(ba[i].Symbol, ticker) {
			if !strings.Contains(ba[i].Symbol, "UP"+ticker) && !strings.Contains(ba[i].Symbol, "DOWN"+ticker) || strings.EqualFold(ba[i].Symbol, "UP"+ticker) || strings.EqualFold(ba[i].Symbol, "DOWN"+ticker) {
				ba[i].Symbol = strings.Replace(ba[i].Symbol, ticker, "", -1)
				if ba[i].Symbol != "" {
					ba[i].Price, _ = strconv.ParseFloat(ba[i].BidPrice, 64)
					result = append(result, ba[i])
				}
			}
		}
	}
	return result
}
