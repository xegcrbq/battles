package binance

import (
	"battles/internal/utils/custom_errs"
	"battles/internal/utils/logger"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
)

type BinanceHolder struct {
	url          string
	dataMap      map[string]float64
	filterTicker string
}

func NewBinanceHolder(url, filterTicker string) *BinanceHolder {
	bh := &BinanceHolder{
		url:          url,
		dataMap:      make(map[string]float64),
		filterTicker: filterTicker,
	}
	bh.init()
	bh.AutoUpdate(30 * time.Second)
	return bh
}
func (h *BinanceHolder) init() {
	data, err := h.getData()
	var ba []binanceAnswer
	err = json.Unmarshal(data, &ba)
	if err != nil {
		logger.Get().Warnf("error unmarshalling binance data: %v", err)
	}
	ba = filterAnswers(ba, h.filterTicker)
	for i := range ba {
		h.dataMap[ba[i].Symbol] = ba[i].Price
	}
	delete(h.dataMap, "")
}

func (h *BinanceHolder) getData() ([]byte, error) {
	agent := fiber.Get(h.url)
	if err := agent.Parse(); err != nil {
		logger.Get().Warnf("error BinanceHolder getData err: %v", err)
		return nil, err
	}
	code, body, errs := agent.Bytes()
	if code == 200 && len(errs) == 0 {
		return body, nil
	}
	logger.Get().Warnf("error BinanceHolder getData code: %v errors: %v", code, errs)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return nil, custom_errs.ParsingError
}

func (h *BinanceHolder) GetPriceByTicker(key string) (float64, bool) {
	res, ok := h.dataMap[key]
	return res, ok
}
func (h *BinanceHolder) GetMap() *map[string]float64 {
	return &h.dataMap
}

// AutoUpdate starts endlessly updating data
func (h *BinanceHolder) AutoUpdate(period time.Duration) {
	if period < time.Second*3 {
		period = time.Second * 10
	}
	go func() {
		for range time.Tick(period) {
			//logger.Get().Debug("Auto Update ", h.dataMap)
			h.init()
		}
	}()
}
