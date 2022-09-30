package balance

import (
	"battles/internal/utils/logger"
	"fmt"
	"testing"
)

func TestHolder(t *testing.T) {
	binance := NewUpdateService(`https://api.binance.com/api/v3/avgPrice`)
	h := NewHolder(binance)
	h.InitTop10()
	lg := logger.Get()
	fmt.Printf("%p\n", lg)
	lg.Infof("pre updated pairs: %v", h.Pairs)
	h.Update()
	lg.Infof("updated pairs: %v", h.Pairs)
}
