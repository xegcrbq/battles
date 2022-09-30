package auth

import (
	"battles/internal/users/user_models"
	"battles/internal/utils/registry"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthController struct {
	reg *registry.Registry
}

func NewAuthController() *AuthController {
	return &AuthController{reg: registry.Get()}
}

func (c *AuthController) Auth(ctx *fiber.Ctx) error {
	pubAdress := ctx.Query("publicaddress")
	c.reg.Log.Debug("Auth with public_address: ", pubAdress)
	paCookie := c.reg.Tknz.NewJWTCookie("public_address_token", pubAdress, time.Now().Add(time.Hour*24*30))
	ctx.Cookie(paCookie)
	return ctx.SendString(fmt.Sprintf("public_address_token created for address: %v", pubAdress))
}
func (c *AuthController) Portfolio(ctx *fiber.Ctx) error {
	paCookie := ctx.Cookies("public_address_token")
	dc, tkn, err := c.reg.Tknz.ParseDataClaims(paCookie)
	if err != nil || !tkn.Valid {
		c.reg.Log.Debug("Portfolio, invalid token")
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	c.reg.Log.Debug("Portfolio, access with address: ", dc.Data)
	answ := c.reg.Repo.Exec(user_models.QueryUserReadByUserPublicAddress{PublicAddress: dc.Data})
	if answ.Err != nil || answ.User == nil {
		c.reg.Log.Warnf("Portfolio, User not found, err: %v", answ.Err)
		return answ.Err
	}
	return ctx.Render(
		"portfolio",
		fiber.Map{"PublicAddress": map[string]string{"BTC": "234", "LTC": "32"}},
	)
}
