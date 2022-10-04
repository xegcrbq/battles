package web_socket

import (
	"battles/internal/base_balance/base_balance_models"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/utils/converter"
	"battles/internal/utils/logger"
	"battles/internal/utils/registry"
	"github.com/antoniodipinto/ikisocket"
	"strconv"
)

func execWSReq(wsReq *WSReq, ep *ikisocket.EventPayload) {
	userId, err := strconv.Atoi(ep.SocketAttributes["UserId"].(string))
	if err != nil {
		logger.Get().Warnf("wsReq strconv.Atoi err  %v", err)
		return
	}
	switch wsReq.ReqType {
	case "buy":
		buyWSReq(&wsReq.UserBalanceCoins, ep, int64(userId))
	default:
		logger.Get().Debugf("wsReq Unknown command '%v'", wsReq.ReqType)
	}
}
func buyWSReq(UserBalanceCoins *WSUserBalanceCoins, ep *ikisocket.EventPayload, userId int64) {
	logger.Get().Debugf("wsReq buy")
	price, found := registry.Get().BinanceHolder.GetPriceByTicker(UserBalanceCoins.Ticker)
	if !found {
		logger.Get().Warnf("wsReq ticker(%v) not found", UserBalanceCoins.Ticker)
		return
	}
	if price <= 0 {
		logger.Get().Warnf("wsReq incorrect price(%v) for %v", price, UserBalanceCoins.Ticker)
		return
	}

	answbb := registry.Get().Repo.Exec(base_balance_models.QueryReadBaseBalanceByUserId{UserId: userId})
	if answbb.Err != nil {
		logger.Get().Warnf("wsReq QueryReadBaseBalanceByUserId err  %v", answbb.Err)
		return
	}
	if answbb.BaseBalance == nil {
		logger.Get().Warnf("wsReq QueryReadBaseBalanceByUserId basebalance not found")
		return
	}
	logger.Get().Debugf("balance before update %v for user %v", answbb.BaseBalance.Amount, userId)
	parsedAmo, err := strconv.ParseFloat(UserBalanceCoins.Amount, 64)
	logger.Get().Debugf("parsedAmo %v", parsedAmo)
	if err != nil {
		logger.Get().Warnf("wsReq strconv.ParseFloat error %v", err)
		return
	}
	if parsedAmo > 2<<33-1 {
		logger.Get().Warnf("wsReq parsedAmo too large %v", parsedAmo)
		return
	}
	coinAmo := converter.Float64ToInt64(parsedAmo)
	usdtAmo := converter.Float64ToInt64(parsedAmo * price)
	logger.Get().Debugf("coinAmo: %v\nusdtAmo: %v", coinAmo, usdtAmo)
	logger.Get().Debugf("answbb.BaseBalance.Amount: %v", answbb.BaseBalance.Amount)
	if answbb.BaseBalance.Amount < usdtAmo {
		logger.Get().Infof("wsReq попытка закупиться на недостающий баланс")
		// !!! нужно отправлять ответ недостаточный баланс
		return
	}

	//желательно делать через транзакции, а не через 2 запроса

	answbb = registry.Get().Repo.Exec(base_balance_models.CommandUpdateBaseBalanceByBaseBalance{BaseBalance: &base_balance_models.BaseBalance{
		UserId: int64(userId),
		Amount: answbb.BaseBalance.Amount - usdtAmo,
	}})

	logger.Get().Debugf("CommandUpdateBaseBalanceByBaseBalance err: %v", answbb.Err)

	answUBC := registry.Get().Repo.Exec(userbalancecoins_models.CommandUpdateOrCreateBalanceByUserIdAmountSpentAndTicker{
		UserId: int64(userId),
		Amount: coinAmo,
		Spent:  usdtAmo,
		Ticker: UserBalanceCoins.Ticker,
	})
	logger.Get().Debugf("CommandUpdateOrCreateBalanceByUserIdAmountSpentAndTicker err: %v", answUBC.Err)
}
