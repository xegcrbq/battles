package portfolio

import (
	"battles/internal/base_balance/base_balance_models"
	"battles/internal/coins/coins_model"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/userbalancecoins/userbalancecoins_models/ubc_convert"
	"battles/internal/users/user_models"
	"battles/internal/utils/converter"
	"battles/internal/utils/logger"
	"battles/internal/utils/registry"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type PortfolioController struct {
	reg *registry.Registry
}

func NewPortfolioController() *PortfolioController {
	InitCoinsDB()
	return &PortfolioController{reg: registry.Get()}
}

func (c *PortfolioController) Portfolio(ctx *fiber.Ctx) error {
	paCookie := ctx.Cookies("public_address_token")
	dc, tkn, err := c.reg.Tknz.ParseDataClaims(paCookie)
	if err != nil || !tkn.Valid {
		c.reg.Log.Debug("Portfolio, invalid token")
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	c.reg.Log.Debug("Portfolio, access with address: ", dc.Data)
	answU := c.reg.Repo.Exec(user_models.QueryUserReadByUserPublicAddress{PublicAddress: dc.Data})
	if answU.Err != nil || answU.User == nil {
		c.reg.Log.Warnf("Portfolio, User not found, err: %v", answU.Err)
		return answU.Err
	}
	answB := c.reg.Repo.Exec(userbalancecoins_models.QueryReadUserBalanceAllCoinsByUserPublicAddress{UserPublicAddress: dc.Data})
	if answB.Err != nil || answB.UserBalanceCoins == nil {
		c.reg.Log.Infof("Portfolio, user balance coin not found, err: %v", answB.Err)
		return answB.Err
	}
	c.reg.Log.Debugf("Balances: %v", answB.UserBalanceCoins)
	answBB := c.reg.Repo.Exec(base_balance_models.QueryReadBaseBalanceByUserId{answU.User.UserId})
	if answBB.Err != nil || answBB.BaseBalance == nil {
		c.reg.Log.Infof("Portfolio, base balance not found, err: %v", answB.Err)
		return answB.Err
	}
	ubcConverted, err := ubc_convert.ConvertUBCarrToUBCwParr(*answB.UserBalanceCoins)
	if err != nil {
		return err
	}
	return ctx.Render(
		"layouts/portfolio",
		fiber.Map{
			"PublicAddress": answU.User.PublicAddress,
			"Balance":       ubcConverted,
			"BaseBalance":   fmt.Sprintf("%.8f", converter.Int64ToFloat64(answBB.BaseBalance.Amount)),
		},
	)
}

func InitCoinsDB() {
	balanceMap := registry.Get().BinanceHolder.GetMap()
	answ := registry.Get().Repo.Exec(coins_model.QueryReadCoinsCount{})
	if answ.Err != nil {
		logger.Get().Debug("PortfolioController InitCoinsDB error: ", answ.Err)
		return
	}
	if len(*balanceMap) > answ.Count {
		logger.Get().Info("Starting coins table initializing")
		for k := range *balanceMap {
			if k != "" {
				registry.Get().Repo.Exec(coins_model.CommandCreateCoinByTicker{Ticker: k})
			}
		}
	}
}
