package main

import (
	"go.uber.org/zap"
	"github.com/spf13/viper"
	"github.com/MrDavudov/TestWB/internal/handler"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := initConfig(); err != nil {
		sugar.Fatalf("error Initializing configs: %s", err)
	}

	if err := handler.Start(viper.GetString("port")); err != nil {
		sugar.Fatalf("errors occured while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}