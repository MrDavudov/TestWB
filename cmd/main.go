package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MrDavudov/TestWB/internal/handler"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.TextFormatter))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Инициализация config.yaml
	if err := initConfig(); err != nil {
		logrus.Fatalf("error Initializing configs: %s", err)
	}

	// Подключения .env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err)
	}

	// Запуск сервера
	srv := new(handler.Server)
	go func() {
		if err := srv.Start(viper.GetString("port")); err != nil {
			logrus.Fatalf("errors occured while running http server: %s", err)
		}
	}()


	logrus.Info("Start server")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			logrus.Errorf("error occured on server shutting down: %s", err.Error())
		}
		logrus.Info("Stop server")
		time.Sleep(5 * time.Second)
	}()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

