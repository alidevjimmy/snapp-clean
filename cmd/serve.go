package cmd

import (
	"fmt"
	"github.com/alidevjimmy/snapp-clean/internal/config"
	"github.com/alidevjimmy/snapp-clean/internal/pkg/logger/zap"
	"github.com/alidevjimmy/snapp-clean/internal/repository/postgres"
	"github.com/alidevjimmy/snapp-clean/internal/transport/http/echo"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"
)

var (
	serveCMD = &cli.Command{
		Name: "serve",
		Aliases: []string{"s"},
		Usage: "serve http",
		Action: serve,
	}
)

func serve(ctx *cli.Context) error {
	cfg := new(config.Config)
	if err := config.ReadYaml("config.yaml", cfg); err != nil {
		panic(err)
	}

	f , err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	logger := zap.New(f, zapcore.ErrorLevel)

	postgresRepo, err := postgres.New(cfg.Postgres, logger)
	if err != nil {
		panic(err)
	}
	_ = postgresRepo

	// send postgresRepo & logger to userService

	restServer := echo.New()
	go func() {
		if err := restServer.Start(cfg.App.Address); err != nil {
			logger.Error(fmt.Sprintf("error while serving: %v", err))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT)
	<-signalChan
	fmt.Println("\nReceived an interrupt, closing connections...")
	if err := restServer.ShutDown(); err != nil {
		fmt.Println("\nRest server doesn't shutdown in 10 seconds")
	}
	return nil
}