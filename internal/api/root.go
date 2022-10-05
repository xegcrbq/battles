package api

import (
	"battles/internal/api/auth"
	"battles/internal/api/portfolio"
	"battles/internal/api/web_socket"
	"battles/internal/utils/logger"
	"context"
	"github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/sirupsen/logrus"
)

type APIService struct {
	log    *logrus.Entry
	router *fiber.App
}

func (svc *APIService) Serve(addr string) {
	svc.log.Fatal(svc.router.Listen(addr))
}

func (svc *APIService) Shutdown(ctx context.Context) error {
	return svc.router.Shutdown()
}

func NewAPIService() (*APIService, error) {
	engine := html.New("./views", ".html")

	svc := &APIService{
		log: logrus.NewEntry(logger.Get()),
		router: fiber.New(fiber.Config{
			Views: engine,
		}),
	}
	authCtrl := auth.NewAuthController()
	portfCtrl := portfolio.NewPortfolioController()
	wbCtrl := web_socket.NewWSController()
	//svc.router.Use(svc.AuthMiddleware())
	api := svc.router.Group("/api/")
	auth := svc.router.Group("/auth/")
	api.Static("/", "./styles")
	//portfolio
	api.Get("portfolio/", portfCtrl.Portfolio)
	//login2
	api.Get("login", authCtrl.Auth)
	api.Get("login2", authCtrl.Auth2)
	//test
	api.Post("test", authCtrl.Test)
	//login with MM
	auth.Post("getNonce", authCtrl.GetNonce)
	//web socket
	api.Get("ws/:public_address_token", ikisocket.New(wbCtrl.SocketReaderCreate))

	return svc, nil
}
