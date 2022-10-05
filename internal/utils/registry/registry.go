package registry

import (
	"battles/internal/balance/binance"
	"battles/internal/utils/logger"
	"battles/internal/utils/repository"
	"battles/internal/utils/tokenizer"
	"github.com/sirupsen/logrus"
	"sync"
)

var reg *Registry = nil
var once sync.Once

type Registry struct {
	Log           *logrus.Logger
	Tknz          *tokenizer.Tokenizer
	Repo          *repository.Repository
	BinanceHolder *binance.BinanceHolder
}

func Get() *Registry {
	once.Do(func() {
		reg = &Registry{
			Log:           logger.Get(),
			Tknz:          tokenizer.Get(),
			Repo:          repository.NewRepository(),
			BinanceHolder: binance.NewBinanceHolder("https://api.binance.com/api/v3/ticker/bookTicker", "USDT"),
		}
	})
	return reg
}
