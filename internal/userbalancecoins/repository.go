package userbalancecoins

import (
	"battles/internal/answer"
	"battles/internal/userbalancecoins/userbalancecoins_models"
)

type UserBalanceCoinsRepo interface {
	ReadUserBalanceCoinsByUserPublicAddress(query userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress) *answer.Answer
}
