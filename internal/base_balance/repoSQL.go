package base_balance

import (
	"battles/internal/answer"
	"battles/internal/base_balance/base_balance_models"
	"battles/internal/utils/logger"
	"github.com/jmoiron/sqlx"
)

type BaseBalanceRepoSQL struct {
	db *sqlx.DB
}

func NewBaseBalanceRepoSQL(db *sqlx.DB) BaseBalanceRepo {
	return &BaseBalanceRepoSQL{db: db}
}

func (r *BaseBalanceRepoSQL) CreateBaseBalanceByBaseBalance(command base_balance_models.CommandCreateBaseBalanceByBaseBalance) *answer.Answer {
	logger.Get().Debug("Trying create BaseBalance with command: ", command.BaseBalance)
	_, err := r.db.Exec(`INSERT INTO base_balances(userid, amount) VALUES ($1, $2)`, command.BaseBalance.UserId, command.BaseBalance.Amount)
	return &answer.Answer{Err: err}
}
func (r *BaseBalanceRepoSQL) ReadBaseBalanceByUserId(query base_balance_models.QueryReadBaseBalanceByUserId) *answer.Answer {
	logger.Get().Debug("Read BaseBalancesByUserId with query: ", query)
	var baseBalance base_balance_models.BaseBalance
	err := r.db.Get(&baseBalance, `SELECT * FROM base_balances WHERE userid = $1`, query.UserId)
	logger.Get().Debug("Read Base balances: ", baseBalance, " for userId: ", query.UserId)
	return &answer.Answer{Err: err, BaseBalance: &baseBalance}
}
