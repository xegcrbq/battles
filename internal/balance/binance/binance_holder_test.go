package binance

import (
	"battles/internal/utils/logger"
	"testing"
)

func TestNewBinanceHolder(t *testing.T) {
	bh := NewBinanceHolder("https://api.binance.com/api/v3/ticker/bookTicker", "USDT")
	logger.Get().Debug(bh.GetMap())
}
