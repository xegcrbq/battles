package userbalancecoins

import (
	"battles/internal/answer"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"github.com/jmoiron/sqlx"
)

type UserBalanceCoinsRepoSQL struct {
	db *sqlx.DB
}

func NewUserBalanceCoinsRepoSQL(db *sqlx.DB) UserBalanceCoinsRepo {
	return &UserBalanceCoinsRepoSQL{db: db}
}

func (r *UserBalanceCoinsRepoSQL) ReadUserBalanceCoinsByUserPublicAddress(query userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress) *answer.Answer {
	return &answer.Answer{Err: nil}
}
