package ubc_convert

import (
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/utils/errors_custom"
	"battles/internal/utils/registry"
	"fmt"
	"math"
)

func UBCToWithPrice(ubc *userbalancecoins_models.UserBalanceCoins) (*userbalancecoins_models.UserBalanceCoinsWithPrice, error) {

	if ubc.Amount/int64(int(math.Pow10(8)))/int64(math.Floor(math.MaxFloat64)) > 0 {
		return nil, errors_custom.VariableTooLarge
	}
	price, err := registry.Get().BalanceHolder.GetByKeyUSDT(ubc.Ticker)
	if err != nil {
		return nil, err
	}
	return &userbalancecoins_models.UserBalanceCoinsWithPrice{
		Ticker: ubc.Ticker,
		Amount: fmt.Sprintf("%.8f", float64(ubc.Amount%int64(math.Pow10(8)))*math.Pow10(-8)+float64(ubc.Amount/int64(int(math.Pow10(8))))),
		Price:  fmt.Sprintf("%.8f", price),
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
