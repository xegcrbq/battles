package repository

import (
	"battles/internal/answer"
	"battles/internal/balance"
	"battles/internal/balance/balance_models"
	"battles/internal/base_balance"
	"battles/internal/base_balance/base_balance_models"
	"battles/internal/buy_history"
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
	bbr  base_balance.BaseBalanceRepo
	bhr  buy_history.BuyHistoryRepo
}

func NewRepository() *Repository {
	return &Repository{
		ur:   users.NewUserRepoSQL(db.Get()),
		br:   balance.NewBalanceRepoSQL(db.Get()),
		ubcr: userbalancecoins.NewUserBalanceCoinsRepoSQL(db.Get()),
		bbr:  base_balance.NewBaseBalanceRepoSQL(db.Get()),
		bhr:  buy_history.NewBuyHistoryRepoSQL(db.Get()),
	}
}

func (r *Repository) Exec(command interface{}) *answer.Answer {
	switch command.(type) {
	//user
	case user_models.CommandUserCreateByUser:
		return r.ur.CreateUserByUser(command.(user_models.CommandUserCreateByUser))
	case user_models.QueryUserReadByUserPublicAddress:
		return r.ur.ReadUserByUserPublicAddress(command.(user_models.QueryUserReadByUserPublicAddress))
	//balance
	case balance_models.QueryReadBalanceByUserIdAndCoinId:
		return r.br.ReadBalancesByUserIdAndCoinId(command.(balance_models.QueryReadBalanceByUserIdAndCoinId))
	case balance_models.QueryReadBalancesByUserId:
		return r.br.ReadBalancesByUserId(command.(balance_models.QueryReadBalancesByUserId))
	case balance_models.CommandCreateBalanceByBalance:
		return r.br.CreateBalanceByBalance(command.(balance_models.CommandCreateBalanceByBalance))
	//user balance coins
	case userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress:
		return r.ubcr.ReadUserBalanceCoinsByUserPublicAddress(command.(userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress))
	case userbalancecoins_models.QueryReadUserBalanceAllCoinsByUserPublicAddress:
		return r.ubcr.ReadUserBalanceAllCoinsByUserPublicAddress(command.(userbalancecoins_models.QueryReadUserBalanceAllCoinsByUserPublicAddress))
	case userbalancecoins_models.CommandUpdateOrCreateBalanceByUserIdAmountAndTicker:
		return r.ubcr.UpdateOrCreateBalanceByUserIdAmountAndTicker(command.(userbalancecoins_models.CommandUpdateOrCreateBalanceByUserIdAmountAndTicker))
	//base balance
	case base_balance_models.CommandCreateBaseBalanceByBaseBalance:
		return r.bbr.CreateBaseBalanceByBaseBalance(command.(base_balance_models.CommandCreateBaseBalanceByBaseBalance))
	case base_balance_models.QueryReadBaseBalanceByUserId:
		return r.bbr.ReadBaseBalanceByUserId(command.(base_balance_models.QueryReadBaseBalanceByUserId))
	case base_balance_models.CommandUpdateBaseBalanceByBaseBalance:
		return r.bbr.UpdateBaseBalanceByBaseBalance(command.(base_balance_models.CommandUpdateBaseBalanceByBaseBalance))
	//buy history
	case buy_history.QueryReadBuyHistorySimpleByUserId:
		return r.bhr.ReadBuyHistorySimpleByUserId(command.(buy_history.QueryReadBuyHistorySimpleByUserId))
	case buy_history.CommandCreateBuyHistoryByBuyHistory:
		return r.bhr.CreateBuyHistoryByBuyHistory(command.(buy_history.CommandCreateBuyHistoryByBuyHistory))
	default:
		return &answer.Answer{
			Err: errors_custom.CommandNotFound,
		}
	}
}
