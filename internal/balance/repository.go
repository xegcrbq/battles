package balance

import (
	"battles/internal/answer"
	"battles/internal/balance/balance_models"
)

type BalanceRepo interface {
	CreateBalanceByBalance(command balance_models.CommandCreateBalanceByBalance) *answer.Answer
	ReadBalanceByUserId(command balance_models.QueryReadBalanceByUserId) *answer.Answer
}
