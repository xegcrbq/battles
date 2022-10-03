package balance

import (
	"battles/internal/utils/logger"
	"testing"
	"time"
)

func TestHolder(t *testing.T) {
	binance := NewUpdateService(`https://api.binance.com/api/v3/avgPrice`)
	h := NewHolder(binance)
	h.InitTop10()
	lg := logger.Get()
	lg.Infof("pre updated pairs: %v", h.Pairs)
	h.Update()
	lg.Infof("updated pairs: %v", h.Pairs)
	price, err := h.GetPriceByKeyUSDT("BTC")
	lg.Infof("BTC price: %v, error: %v", price, err)
	h.AutoUpdate(time.Second * 5)
}
