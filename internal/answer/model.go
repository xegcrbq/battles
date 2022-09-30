package answer

import (
	"battles/internal/balance/balance_models"
	"battles/internal/coins"
	"battles/internal/users/user_models"
)

type Answer struct {
	Err     error
	User    *user_models.User
	Coin    *coins.Coin
	Balance *balance_models.ClientBalance
}
