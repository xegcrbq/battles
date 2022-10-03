package answer

import (
	"battles/internal/balance/balance_models"
	"battles/internal/base_balance/base_balance_models"
	"battles/internal/coins"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/users/user_models"
)

type Answer struct {
	Err              error
	User             *user_models.User
	Coin             *coins.Coin
	Balance          *balance_models.Balance
	Balances         *[]balance_models.Balance
	BaseBalance      *base_balance_models.BaseBalance
	UserBalanceCoins *[]userbalancecoins_models.UserBalanceCoins
	//BuyHistorySums key - ticker, value - grouped by coinId sum
	BuyHistorySums *map[string]string
}
