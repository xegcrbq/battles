package ubc_convert

import (
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/utils/converter"
	"battles/internal/utils/custom_errs"
	"battles/internal/utils/logger"
	"battles/internal/utils/registry"
	"fmt"
	"math"
)

func UBCToWithPrice(ubc *userbalancecoins_models.UserBalanceCoins) (*userbalancecoins_models.UserBalanceCoinsWithPrice, error) {

	if ubc.Amount/int64(math.Pow10(8))/int64(math.Floor(math.MaxFloat64)) > 0 {
		return nil, custom_errs.VariableTooLarge
	}
	price, success := registry.Get().BinanceHolder.GetPriceByTicker(ubc.Ticker)
	if !success {
		logger.Get().Warnf("UBCToWithPrice Ticker '%v' not found", ubc.Ticker)
		return nil, custom_errs.VariableNotFound
	}
	return &userbalancecoins_models.UserBalanceCoinsWithPrice{
		Ticker: ubc.Ticker,
		Amount: fmt.Sprintf("%.8f", converter.Int64ToFloat64(ubc.Amount)),
		Price:  fmt.Sprintf("%.8f", price),
		Spent:  fmt.Sprintf("%.8f", converter.Int64ToFloat64(ubc.Spent)),
	}, nil
}
func ConvertUBCarrToUBCwParr(ubc []userbalancecoins_models.UserBalanceCoins) ([]userbalancecoins_models.UserBalanceCoinsWithPrice, error) {
	result := make([]userbalancecoins_models.UserBalanceCoinsWithPrice, len(ubc))
	for i := range ubc {
		val, err := UBCToWithPrice(&ubc[i])
		if err != nil {
			return nil, err
		}
		result[i] = *val
	}
	return result, nil
}
