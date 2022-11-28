package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/MrDavudov/TestWB/internal/handler"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.TextFormatter))

	// Инициализация config.yaml
	if err := initConfig(); err != nil {
		logrus.Fatalf("error Initializing configs: %s", err)
	}

	// Подключения .env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err)
	}

	srv := new(handler.Server)
	go func() {
		if err := srv.Start(viper.GetString("port")); err != nil {
			logrus.Fatalf("errors occured while running http server: %s", err)
		}
	}()

	// // ассинхроонное обновление температуры каждую минуту
	// go func() {
	// 	for {
	// 		if err := service.SaveAsync(); err != nil {
	// 			logrus.Fatalf("failed save async in db: %s", err)
	// 		}
	// 	}
	// }()

	logrus.Info("Start server")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("Shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}


}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

