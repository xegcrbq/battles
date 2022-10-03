package portfolio

import (
	"battles/internal/base_balance/base_balance_models"
	"battles/internal/buy_history"
	"battles/internal/userbalancecoins/userbalancecoins_models"
	"battles/internal/userbalancecoins/userbalancecoins_models/ubc_convert"
	"battles/internal/users/user_models"
	"battles/internal/utils/registry"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math"
)

type PortfolioController struct {
	reg *registry.Registry
}

func NewPortfolioController() *PortfolioController {
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
	answBH := c.reg.Repo.Exec(buy_history.QueryReadBuyHistorySimpleByUserId{UserId: answU.User.UserId})
	if answBH.Err != nil || answBH.BuyHistorySums == nil {
		c.reg.Log.Infof("Portfolio, BuyHistory not found, err: %v", answBH.Err)
		return answBH.Err
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
			"BaseBalance":   fmt.Sprintf("%.8f", float64(answBB.BaseBalance.Amount)/math.Pow10(8)),
			"BuyHistory":    answBH.BuyHistorySums,
		},
	)
}
