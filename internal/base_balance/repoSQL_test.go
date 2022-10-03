package base_balance

import (
	"battles/internal/base_balance/base_balance_models"
	"battles/internal/db"
	"battles/internal/utils/logger"
	"testing"
)

func TestBaseBalanceRepo(t *testing.T) {
	bbr := NewBaseBalanceRepoSQL(db.Get())
	answ := bbr.ReadBaseBalanceByUserId(base_balance_models.QueryReadBaseBalanceByUserId{UserId: 1})
	logger.Get().Debugf("ReadBaseBalanceByUserId value: %v err: %v", answ.BaseBalance, answ.Err)

	answC := bbr.CreateBaseBalanceByBaseBalance(base_balance_models.CommandCreateBaseBalanceByBaseBalance{BaseBalance: &base_balance_models.BaseBalance{
		UserId: 2,
		Amount: 1000 * 10000000,
	}})
	logger.Get().Debugf("CreateBaseBalanceByBaseBalance err: %v", answC.Err)
}
