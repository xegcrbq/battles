package balance

import (
	"battles/internal/balance/balance_models"
	"battles/internal/utils/logger"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
)

type Holder struct {
	Pairs         []balance_models.Pair
	UpdateService *UpdateService
}

func NewHolder(us *UpdateService) *Holder {
	return &Holder{UpdateService: us, Pairs: []balance_models.Pair{}}
}
func (h *Holder) InitTop10() {
	h.AddPair(`BTC`, `USDT`)
	h.AddPair(`ETH`, `USDT`)
	h.AddPair(`USDC`, `USDT`)
	h.AddPair(`BNB`, `USDT`)
	h.AddPair(`XRP`, `USDT`)
	h.AddPair(`BUSD`, `USDT`)
	h.AddPair(`ADA`, `USDT`)
	h.AddPair(`SOL`, `USDT`)
	h.AddPair(`DOGE`, `USDT`)
	h.AddPair(`DOT`, `USDT`)
}

func (h *Holder) AddPair(first, second string) {
	h.Pairs = append(h.Pairs, balance_models.Pair{
		First:  first,
		Second: second,
	})
}
func (h *Holder) Update() {
	lg := logger.Get()
	wg := sync.WaitGroup{}
	wg.Add(len(h.Pairs))
	fmt.Printf("%p\n", lg)
	for index, _ := range h.Pairs {
		go func(wg *sync.WaitGroup, i int) {
			a := binanceAnswer{}
			data := h.UpdateService.GetWithSymbol(h.Pairs[i].String())
			err := json.Unmarshal(data, &a)
			if err == nil {
				//lg.Debugf("a: %v", a.Price)
				newPrice, err2 := strconv.ParseFloat(a.Price, 64)
				//lg.Debugf("a: %v", newPrice)
				if err2 == nil {
					h.Pairs[i].Price = newPrice
				}
			}
			wg.Done()
		}(&wg, index)
	}
	wg.Wait()
}