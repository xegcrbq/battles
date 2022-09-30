package balance

import (
	"battles/internal/answer"
	"battles/internal/balance/balance_models"
	"battles/internal/utils/logger"
	"github.com/jmoiron/sqlx"
)

type BalanceRepoSQL struct {
	db *sqlx.DB
}

func NewBalanceRepoSQL(db *sqlx.DB) BalanceRepo {
	return &BalanceRepoSQL{db: db}
}

func (r *BalanceRepoSQL) CreateBalanceByBalance(command balance_models.CommandCreateBalanceByBalance) *answer.Answer {
	b := command.Balance
	logger.Get().Debug("Created Balance with command: ", command.Balance)
	_, err := r.db.Exec(`INSERT INTO balances(userid, amount, coinid) VALUES ($1, $2, $3)`, b.UserId, b.Amount, b.CoinId)
	return &answer.Answer{Err: err}
}
func (r *BalanceRepoSQL) ReadBalancesByUserId(query balance_models.QueryReadBalancesByUserId) *answer.Answer {
	logger.Get().Debug("ReadBalancesByUserId with query: ", query)
	var balances []balance_models.Balance
	err := r.db.Select(&balances, `SELECT * FROM balances WHERE userid = $1`, query.UserId)
	logger.Get().Debug("Read balances: ", balances, " for userId: ", query.UserId)
	return &answer.Answer{Err: err, Balances: &balances}
}
func (r *BalanceRepoSQL) ReadBalancesByUserIdAndCoinId(query balance_models.QueryReadBalanceByUserIdAndCoinId) *answer.Answer {
	logger.Get().Debug("ReadBalancesByUserIdAndCoinId with query: ", query)
	var balance balance_models.Balance
	err := r.db.Get(&balance, `SELECT * FROM balances WHERE userid = $1 and coinid = $2`, query.UserId, query.CoinId)
	logger.Get().Debugf("Read balance: %v for userId: %v and coinId: %v ", balance, query.UserId, query.CoinId)
	return &answer.Answer{Err: err, Balance: &balance}
}
