package balance

import (
	"battles/internal/answer"
	"battles/internal/balance/balance_models"
)

type BalanceRepo interface {
	CreateBalanceByBalance(command balance_models.CommandCreateBalanceByBalance) *answer.Answer
	ReadBalancesByUserId(command balance_models.QueryReadBalancesByUserId) *answer.Answer
	ReadBalancesByUserIdAndCoinId(query balance_models.QueryReadBalanceByUserIdAndCoinId) *answer.Answer
}
