package repository

import (
	"battles/internal/answer"
	"battles/internal/balance"
	"battles/internal/balance/balance_models"
	"battles/internal/db"
	"battles/internal/userbalancecoins"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/users"
	"battles/internal/users/user_models"
	"battles/internal/utils/errors_custom"
)

type Repository struct {
	ur   users.UserRepo
	br   balance.BalanceRepo
	ubcr userbalancecoins.UserBalanceCoinsRepo
}

func NewRepository() *Repository {
	return &Repository{
		ur:   users.NewUserRepoSQL(db.Get()),
		br:   balance.NewBalanceRepoSQL(db.Get()),
		ubcr: userbalancecoins.NewUserBalanceCoinsRepoSQL(db.Get()),
	}
}

func (r *Repository) Exec(command interface{}) *answer.Answer {
	switch command.(type) {
	case user_models.CommandUserCreateByUser:
		return r.ur.CreateUserByUser(command.(user_models.CommandUserCreateByUser))
	case user_models.QueryUserReadByUserPublicAddress:
		return r.ur.ReadUserByUserPublicAddress(command.(user_models.QueryUserReadByUserPublicAddress))
	case balance_models.QueryReadBalanceByUserIdAndCoinId:
		return r.br.ReadBalancesByUserIdAndCoinId(command.(balance_models.QueryReadBalanceByUserIdAndCoinId))
	case balance_models.QueryReadBalancesByUserId:
		return r.br.ReadBalancesByUserId(command.(balance_models.QueryReadBalancesByUserId))
	case balance_models.CommandCreateBalanceByBalance:
		return r.br.CreateBalanceByBalance(command.(balance_models.CommandCreateBalanceByBalance))
	case userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress:
		return r.ubcr.ReadUserBalanceCoinsByUserPublicAddress(command.(userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress))
	case userbalancecoins_models.QueryReadUserBalanceAllCoinsByUserPublicAddress:
		return r.ubcr.ReadUserBalanceAllCoinsByUserPublicAddress(command.(userbalancecoins_models.QueryReadUserBalanceAllCoinsByUserPublicAddress))
	default:
		return &answer.Answer{
			Err: errors_custom.CommandNotFound,
		}
	}
}
