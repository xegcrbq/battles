package registry

import (
	"battles/internal/balance"
	"battles/internal/balance/binance"
	"battles/internal/utils/logger"
	"battles/internal/utils/repository"
	"battles/internal/utils/tokenizer"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var reg *Registry = nil
var once sync.Once

type Registry struct {
	Log           *logrus.Logger
	Tknz          *tokenizer.Tokenizer
	Repo          *repository.Repository
	BalanceHolder *balance.Holder
	BinanceHolder *binance.BinanceHolder
}

func Get() *Registry {
	once.Do(func() {
		us := balance.NewUpdateService(`https://api.binance.com/api/v3/avgPrice`)
		h := balance.NewHolder(us)
		h.InitTop10()
		h.Update()
		go h.AutoUpdate(time.Second * 30)
		reg = &Registry{
			Log:           logger.Get(),
			Tknz:          tokenizer.Get(),
			Repo:          repository.NewRepository(),
			BalanceHolder: h,
			BinanceHolder: binance.NewBinanceHolder("https://api.binance.com/api/v3/ticker/bookTicker", "USDT"),
		}
	})
	return reg
}
