package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/MrDavudov/TestWB/internal/handler"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := initConfig(); err != nil {
		sugar.Fatalf("error Initializing configs: %s", err)
	}

	srv := new(handler.Server)
	go func() {
		if err := srv.Start(viper.GetString("port")); err != nil {
			sugar.Fatalf("errors occured while running http server: %s", err)
		}
	}()

	sugar.Info("Start server")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	sugar.Info("Shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		sugar.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}