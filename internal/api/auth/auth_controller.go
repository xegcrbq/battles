package auth

import (
	"battles/internal/utils/logger"
	"battles/internal/utils/registry"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
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
func (c *AuthController) Auth2(ctx *fiber.Ctx) error {
	return ctx.Render(
		"layouts/login",
		fiber.Map{
			"csrf_token": "token",
		})
}
func (c *AuthController) Test(ctx *fiber.Ctx) error {
	fmt.Println(string(ctx.Body()))
	return nil
}
func (c *AuthController) GetNonce(ctx *fiber.Ctx) error {
	address := ctx.Query("address")
	if address == "" {
		return nil
	}
	nonce := 34532453245
	logger.Get().Debugf("address %v executed GetNonce with nonce: %v", address, nonce)
	return ctx.SendString(strconv.Itoa(nonce))
}
