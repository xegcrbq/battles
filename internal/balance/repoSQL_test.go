package balance

import (
	"battles/internal/balance/balance_models"
	"battles/internal/db"
	"battles/internal/utils/logger"
	"testing"
)

func TestRepoSQLCreate(t *testing.T) {
	br := NewBalanceRepoSQL(db.Get())
	answ := br.CreateBalanceByBalance(balance_models.CommandCreateBalanceByBalance{Balance: &balance_models.Balance{UserId: 1, CoinId: 2, Amount: 1000}})
	logger.Get().Debug(answ.Err)
	answ = br.ReadBalancesByUserId(balance_models.QueryReadBalancesByUserId{UserId: 1})
	logger.Get().Debug(answ.Balances, answ.Err)
	answ = br.ReadBalancesByUserIdAndCoinId(balance_models.QueryReadBalanceByUserIdAndCoinId{UserId: 1, CoinId: 2})
	logger.Get().Debug(answ.Balance, answ.Err)
}
