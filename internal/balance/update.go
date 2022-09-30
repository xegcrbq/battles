package balance

import (
	"github.com/gofiber/fiber/v2"
)

type UpdateService struct {
	url string
}
type binanceAnswer struct {
	Price string `json:"price"`
}

func NewUpdateService(url string) *UpdateService {
	return &UpdateService{url: url}
}

func (s *UpdateService) GetWithSymbol(pair string) []byte {
	agent := fiber.Get(s.url + "?symbol=" + pair)
	if err := agent.Parse(); err != nil {
		panic(err)
	}
	code, body, errs := agent.Bytes()
	if code == 200 && len(errs) == 0 {
		return body
	}
	return nil
}
