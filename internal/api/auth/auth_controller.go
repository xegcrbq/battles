package auth

import (
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
