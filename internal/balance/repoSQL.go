package balance

import (
	"battles/internal/answer"
	"battles/internal/balance/balance_models"
	"github.com/jmoiron/sqlx"
)

type BalanceRepoSQL struct {
	db *sqlx.DB
}

func NewBalanceRepoSQL(db *sqlx.DB) BalanceRepo {
	return &BalanceRepoSQL{db: db}
}

func (r *BalanceRepoSQL) CreateBalanceByBalance(command balance_models.CommandCreateBalanceByBalance) *answer.Answer {

	return &answer.Answer{Err: nil}
}
func (r *BalanceRepoSQL) ReadBalanceByUserId(command balance_models.QueryReadBalanceByUserId) *answer.Answer {

	return &answer.Answer{Err: nil}
}

