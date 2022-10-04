package coins

import (
	"battles/internal/answer"
	"battles/internal/coins/coins_model"
)

type CoinsRepo interface {
	CreateCoin(command coins_model.CommandCreateCoinByTicker) *answer.Answer
	GetCoinsCount(query coins_model.QueryReadCoinsCount) *answer.Answer
}
