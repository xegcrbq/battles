package answer

import (
	"battles/internal/balance/balance_models"
	"battles/internal/base_balance/base_balance_models"
	"battles/internal/coins/coins_model"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/users/user_models"
)

type Answer struct {
	Balance     *balance_models.Balance
	Balances    *[]balance_models.Balance
	BaseBalance *base_balance_models.BaseBalance
	//BuyHistorySums key - ticker, value - grouped by coinId sum
	BuyHistorySums   *map[string]string
	Coin             *coins_model.Coin
	Count            int
	Err              error
	UserBalanceCoins *[]userbalancecoins_models.UserBalanceCoins
	User             *user_models.User
}
