package balance

import (
	"testing"
)

func TestGet(t *testing.T) {
	binance := NewUpdateService(`https://api.binance.com/api/v3/avgPrice`)
	binance.GetWithSymbol(`ETHUSDT`)
}
