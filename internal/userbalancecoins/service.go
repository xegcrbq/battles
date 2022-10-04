package userbalancecoins

import (
	"battles/internal/answer"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"github.com/jmoiron/sqlx"
)

type UBCService struct {
	repo UserBalanceCoinsRepo
}
type UBCServiceI interface {
	CreateBalanceByUserPublicAddressAndAmountAndTicker(command userbalancecoins_models.CommandCreateBalanceByUserPublicAddressAndAmountAndTicker) *answer.Answer
	ReadUserBalanceCoinsByUserPublicAddress(query userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress) *answer.Answer
	ReadUserBalanceAllCoinsByUserPublicAddress(query userbalancecoins_models.QueryReadUserBalanceAllCoinsByUserPublicAddress) *answer.Answer
	UpdateOrCreateBalanceByUserIdAmountSpentAndTicker(command userbalancecoins_models.CommandUpdateOrCreateBalanceByUserIdAmountSpentAndTicker) *answer.Answer
}

func NewUBCService(db *sqlx.DB) UBCServiceI {
	return &UBCService{
		repo: NewUserBalanceCoinsRepoSQL(db),
	}
}

func (s *UBCService) CreateBalanceByUserPublicAddressAndAmountAndTicker(command userbalancecoins_models.CommandCreateBalanceByUserPublicAddressAndAmountAndTicker) *answer.Answer {
	return s.repo.CreateBalanceByUserPublicAddressAndAmountAndTicker(&command)
}

func (s *UBCService) ReadUserBalanceCoinsByUserPublicAddress(query userbalancecoins_models.QueryReadUserBalanceCoinsByUserPublicAddress) *answer.Answer {
	return s.repo.ReadUserBalanceCoinsByUserPublicAddress(&query)
}

func (s *UBCService) ReadUserBalanceAllCoinsByUserPublicAddress(query userbalancecoins_models.QueryReadUserBalanceAllCoinsByUserPublicAddress) *answer.Answer {
	return s.repo.ReadUserBalanceAllCoinsByUserPublicAddress(&query)
}
func (s *UBCService) UpdateOrCreateBalanceByUserIdAmountSpentAndTicker(command userbalancecoins_models.CommandUpdateOrCreateBalanceByUserIdAmountSpentAndTicker) *answer.Answer {

	answ := s.repo.ReadCountByUserIdAndTicker(&userbalancecoins_models.QueryReadCountByUserIdAndTicker{
		UserId: command.UserId,
		Ticker: command.Ticker,
	})
	if answ.Err != nil {
		return &answer.Answer{Err: answ.Err}
	}
	if answ.Count == 0 {
		return s.repo.CreateBalanceByUserIdAmountSpentAndTicker(&userbalancecoins_models.CommandCreateBalanceByUserIdAmountSpentAndTicker{
			UserId: command.UserId,
			Amount: command.Amount,
			Spent:  command.Spent,
			Ticker: command.Ticker,
		})
	} else {
		return s.repo.UpdateBalanceByUserIdAmountSpentAndTicker(&userbalancecoins_models.CommandUpdateBalanceByUserIdAmountSpentAndTicker{
			UserId: command.UserId,
			Amount: command.Amount,
			Spent:  command.Spent,
			Ticker: command.Ticker,
		})
	}
}
