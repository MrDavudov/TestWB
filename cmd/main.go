package main

import (
	"github.com/MrDavudov/TestWB/internal/handler"
	"go.uber.org/zap"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	
)

func main() {
	// logrus.SetFormatter(new(logrus.TextFormatter))
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