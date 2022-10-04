package repository

import (
	"battles/internal/answer"
	"battles/internal/balance"
	"battles/internal/balance/balance_models"
	"battles/internal/base_balance"
	"battles/internal/base_balance/base_balance_models"
	"battles/internal/coins"
	"battles/internal/coins/coins_model"
	"battles/internal/db"
	"battles/internal/userbalancecoins"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/users"
	"battles/internal/users/user_models"
	"battles/internal/utils/custom_errs"
	"battles/internal/utils/logger"
)

type Repository struct {
	ur  users.UserRepo
	br  balance.BalanceRepo
	cr  coins.CoinsRepo
	bbr base_balance.BaseBalanceRepo
	ubc userbalancecoins.UBCServiceI
}

func NewRepository() *Repository {
	return &Repository{
		ur:  users.NewUserRepoSQL(db.Get()),
		br:  balance.NewBalanceRepoSQL(db.Get()),
		cr:  coins.NewCoinsRepoSQL(db.Get()),
		bbr: base_balance.NewBaseBalanceRepoSQL(db.Get()),
		ubc: userbalancecoins.NewUBCService(db.Get()),
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
		return r.ubc.ReadUserBalanceCoinsByUserPublicAddress(command.(userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress))
	case userbalancecoins_models.QueryReadUserBalanceAllCoinsByUserPublicAddress:
		return r.ubc.ReadUserBalanceAllCoinsByUserPublicAddress(command.(userbalancecoins_models.QueryReadUserBalanceAllCoinsByUserPublicAddress))
	case userbalancecoins_models.CommandUpdateOrCreateBalanceByUserIdAmountSpentAndTicker:
		return r.ubc.UpdateOrCreateBalanceByUserIdAmountSpentAndTicker(command.(userbalancecoins_models.CommandUpdateOrCreateBalanceByUserIdAmountSpentAndTicker))
	//base balance
	case base_balance_models.CommandCreateBaseBalanceByBaseBalance:
		return r.bbr.CreateBaseBalanceByBaseBalance(command.(base_balance_models.CommandCreateBaseBalanceByBaseBalance))
	case base_balance_models.QueryReadBaseBalanceByUserId:
		return r.bbr.ReadBaseBalanceByUserId(command.(base_balance_models.QueryReadBaseBalanceByUserId))
	case base_balance_models.CommandUpdateBaseBalanceByBaseBalance:
		return r.bbr.UpdateBaseBalanceByBaseBalance(command.(base_balance_models.CommandUpdateBaseBalanceByBaseBalance))
	//coins
	case coins_model.CommandCreateCoinByTicker:
		return r.cr.CreateCoin(command.(coins_model.CommandCreateCoinByTicker))
	case coins_model.QueryReadCoinsCount:
		return r.cr.GetCoinsCount(command.(coins_model.QueryReadCoinsCount))
	default:
		logger.Get().Warnf("Uncknown command sended to repository")
		return &answer.Answer{
			Err: custom_errs.CommandNotFound,
		}
	}
}
