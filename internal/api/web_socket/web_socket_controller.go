package web_socket

import (
	"battles/internal/base_balance/base_balance_models"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/users/user_models"
	"battles/internal/utils/logger"
	"battles/internal/utils/registry"
	"encoding/json"
	"github.com/antoniodipinto/ikisocket"
	"math"
	"strconv"
)

type WSController struct {
	reg *registry.Registry
}

func NewWSController() *WSController {
	c := &WSController{reg: registry.Get()}
	c.Init()
	return c
}

func (c *WSController) Init() {
	ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
		if ep.SocketAttributes["PublicAddress"] != nil && ep.SocketAttributes["UserId"] != nil {
			c.reg.Log.Debugf("Connected: %v", ep.SocketAttributes["PublicAddress"].(string))
			c.reg.Log.Debugf("Connected: %v(UserId)", ep.SocketAttributes["UserId"].(string))
		}
	})

	ikisocket.On(ikisocket.EventMessage, c.wsReq)

	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		if ep.SocketAttributes["PublicAddress"] != nil && ep.SocketAttributes["UserId"] != nil {
			c.reg.Log.Debugf("Disconnected: %v", ep.SocketAttributes["PublicAddress"].(string))
			//c.socketService.DeleteSocket(ep.SocketAttributes["username"].(string), ep.SocketUUID)
		}
	})
}

func (c *WSController) wsReq(ep *ikisocket.EventPayload) {
	// !!! По хорошему это нужно в сервис кидать
	if ep.SocketAttributes["PublicAddress"] == nil || ep.SocketAttributes["UserId"] == nil {
		return
	}
	var wsReq WSReq
	err := json.Unmarshal(ep.Data, &wsReq)
	if err != nil {
		c.reg.Log.Warnf(`Incorrect socket req, Err: %v`, err)
		return
	}
	c.reg.Log.Debugf(`wsReq get data: "%v`, wsReq)
	switch wsReq.ReqType {
	case "buy":
		logger.Get().Debugf("wsReq buy")
		price, err := registry.Get().BalanceHolder.GetPriceByKeyUSDT(wsReq.UserBalanceCoins.Ticker)
		if err != nil {
			logger.Get().Warnf("wsReq ticker(%v) not found", wsReq.UserBalanceCoins.Ticker)
			return
		}
		if price <= 0 {
			logger.Get().Warnf("wsReq incorrect price(%v) for %v", price, wsReq.UserBalanceCoins.Ticker)
			return
		}
		userId, err := strconv.Atoi(ep.SocketAttributes["UserId"].(string))
		if err != nil {
			logger.Get().Warnf("wsReq strconv.Atoi err  %v", err)
			return
		}
		answbb := registry.Get().Repo.Exec(base_balance_models.QueryReadBaseBalanceByUserId{UserId: int64(userId)})
		if answbb.Err != nil {
			logger.Get().Warnf("wsReq QueryReadBaseBalanceByUserId err  %v", err)
			return
		}
		if answbb.BaseBalance == nil {
			logger.Get().Warnf("wsReq QueryReadBaseBalanceByUserId basebalance not found")
			return
		}
		logger.Get().Debugf("balance before update %v for user %v", answbb.BaseBalance.Amount, answbb.BaseBalance.Amount)
		parsedAmo, err := strconv.ParseFloat(wsReq.UserBalanceCoins.Amount, 64)
		parsedAmoInt := int64(math.Floor(parsedAmo * math.Pow10(8)))
		if err != nil {
			logger.Get().Warnf("wsReq strconv.ParseFloat error %v", err)
			return
		}
		if answbb.BaseBalance.Amount < int64(math.Floor(parsedAmo*price*math.Pow10(8))) {
			logger.Get().Infof("wsReq попытка закупиться на недостающий баланс")
			// !!! нужно отправлять ответ недостаточный баланс
			return
		}

		//желательно делать через транзакции, а не через 2 запроса

		answbb = registry.Get().Repo.Exec(base_balance_models.CommandUpdateBaseBalanceByBaseBalance{BaseBalance: &base_balance_models.BaseBalance{
			UserId: int64(userId),
			Amount: answbb.BaseBalance.Amount - int64(math.Floor(parsedAmo*price*math.Pow10(8))),
		}})

		logger.Get().Debugf("CommandUpdateBaseBalanceByBaseBalance err: %v", answbb.Err)
		answUBC := registry.Get().Repo.Exec(userbalancecoins_models.CommandUpdateOrCreateBalanceByUserIdAmountAndTicker{
			UserId: int64(userId),
			Amount: parsedAmoInt,
			Ticker: wsReq.UserBalanceCoins.Ticker,
		})
		logger.Get().Debugf("CommandUpdateBaseBalanceByBaseBalance err: %v", answUBC.Err)
	default:
		logger.Get().Debugf("wsReq Unknown command '%v'", wsReq.ReqType)
	}

}
func (c *WSController) SocketReaderCreate(kws *ikisocket.Websocket) {
	data, tkn, _ := c.reg.Tknz.ParseDataClaims(kws.Params("public_address_token"))
	if !tkn.Valid {
		kws.Close()
		return
	}
	answ := registry.Get().Repo.Exec(user_models.QueryUserReadByUserPublicAddress{PublicAddress: data.Data})
	if answ.Err != nil || answ.User == nil {
		kws.Close()
		return
	}
	kws.SetAttribute("PublicAddress", data.Data)
	logger.Get().Debugf("UesrID: %v", answ.User.UserId)
	kws.SetAttribute("UserId", strconv.Itoa(int(answ.User.UserId)))
	return
}
