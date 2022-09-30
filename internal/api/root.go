package api

import (
	"battles/internal/auth"
	"battles/internal/utils/logger"
	"context"
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
	engine := html.New("./templates", ".html")

	svc := &APIService{
		log: logrus.NewEntry(logger.Get()),
		router: fiber.New(fiber.Config{
			Views: engine,
		}),
	}
	authCtrl := auth.NewAuthController()

	//svc.router.Use(svc.AuthMiddleware())
	api := svc.router.Group("/api/")
	api.Static("/", "./templates")
	api.Get("portfolio/", authCtrl.Portfolio)
	api.Get("login", authCtrl.Auth)

	return svc, nil
}
