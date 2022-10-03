package ubc_convert

import (
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/utils/logger"
	"battles/internal/utils/registry"
	"math"
	"testing"
)

// конвертация не идеальна, но из-за хранения в int потеря точности происходит только при отображении баланса(при создании и обновлении идет работа с int)
func TestConvert(t *testing.T) {
	ubc := userbalancecoins_models.UserBalanceCoins{
		Ticker: "BTC",
		Amount: math.MaxInt64,
	}
	logger.Get().Debugln(math.MaxInt64)
	val, err := UBCToWithPrice(&ubc)
	logger.Get().Debugf("TestConvert value: %.32v, error: %v", val, err)
}
func TestConvert2(t *testing.T) {
	ubc := userbalancecoins_models.UserBalanceCoins{
		Ticker: "BTC",
		Amount: int64(math.Pow10(12)) + 341765193847,
	}
	logger.Get().Debugln(int64(math.Pow10(12)))
	val, err := UBCToWithPrice(&ubc)
	logger.Get().Debugf("TestConvert value: %.32v, error: %v", val, err)
}
func TestConvert3(t *testing.T) {
	answ := registry.Get().Repo.Exec(userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress{UserPublicAddress: "0x8a8cB99FBE417c2fBED13B4982e4fE1BE364d58C"})
	res, err := ConvertUBCarrToUBCwParr(*answ.UserBalanceCoins)
	logger.Get().Debugf("TestConvert3 result: %v\nerror : %v", res, err)

}
