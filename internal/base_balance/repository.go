package base_balance

import (
	"battles/internal/answer"
	"battles/internal/base_balance/base_balance_models"
)

type BaseBalanceRepo interface {
	CreateBaseBalanceByBaseBalance(command base_balance_models.CommandCreateBaseBalanceByBaseBalance) *answer.Answer
	ReadBaseBalanceByUserId(query base_balance_models.QueryReadBaseBalanceByUserId) *answer.Answer
	UpdateBaseBalanceByBaseBalance(command base_balance_models.CommandUpdateBaseBalanceByBaseBalance) *answer.Answer
}
