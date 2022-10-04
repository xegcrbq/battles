package userbalancecoins

import (
	"battles/internal/db"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/utils/logger"
	"testing"
)

func TestRead(t *testing.T) {
	ubcr := NewUserBalanceCoinsRepoSQL(db.Get())
	answ := ubcr.ReadUserBalanceCoinsByUserPublicAddress(
		userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress{
			UserPublicAddress: "0x8a8cB99FBE417c2fBED13B4982e4fE1BE364d58C"})
	logger.Get().Debugf("TestRead userwithbalance: %v \nerror: %v", answ.UserBalanceCoins, answ.Err)
}
func TestCreate(t *testing.T) {
	ubcr := NewUserBalanceCoinsRepoSQL(db.Get())
	answ := ubcr.CreateBalanceByUserPublicAddressAndAmountAndTicker(
		userbalancecoins_models.CommandCreateBalanceByUserPublicAddressAndAmountAndTicker{
			UserPublicAddress: "0x8a8cB99FBE417c2fBED13B4982e4fE1BE364d58C",
			Ticker:            "BNB",
			Amount:            80000,
		})
	logger.Get().Debugf("TestCreate error: %v", answ.Err)
}
func TestUpdateOrCreateBalanceByUserIdAmountAndTicker(t *testing.T) {
	ubcr := NewUserBalanceCoinsRepoSQL(db.Get())
	answ := ubcr.UpdateOrCreateBalanceByUserIdAmountSpentAndTicker(
		userbalancecoins_models.CommandUpdateOrCreateBalanceByUserIdAmountSpentAndTicker{
			UserId: 1,
			Ticker: "DOT",
			Amount: 80500,
		})
	logger.Get().Debugf("TestCreate error: %v", answ.Err)
}
