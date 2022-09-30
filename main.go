package main

import (
	"battles/internal/api"
	"battles/internal/db"
	"battles/internal/utils/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

const (
	defaultAddress   = "0.0.0.0"
	defaultPort      = "8080"
	shutdown_timeout = 5
)

func main() {
	// -------------------- Set up logging -------------------- //

	log := logger.Get()

	// -------------------- Set database -------------------- //
	defer db.Get().Close()
	if err := db.Get().Ping(); err != nil {
		log.Fatalf("failed to check (ping) db connection: %s", err)
		panic(err)
	}
	log.Info("connected to db")

	// -------------------- Set up service -------------------- //

	svc, err := api.NewAPIService()
	if err != nil {
		log.Fatalf("error creating service instance: %s", err)
	}

	go svc.Serve(fmt.Sprintf("%v:%v", defaultAddress, defaultPort))

	// -------------------- Listen for INT signal -------------------- //

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	//fmt.Println("listen")
	<-quit

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*time.Duration(shutdown_timeout),
	)
	defer cancel()

	if err := svc.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("end")
}
