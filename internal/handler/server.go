package handler

import (
	"os"

	"github.com/MrDavudov/TestWB/internal/repository"
	"github.com/MrDavudov/TestWB/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	_ "github.com/lib/pq"
)

func Start(port string) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := initConfig(); err != nil {
		sugar.Fatalf("error Initializing configs: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		sugar.Fatalf("error loading env variables: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		sugar.Fatalf("failed to initialize db: %s", err)
	}
	
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	srv := NewHandler(port, services)

	sugar.Info("Start server")

	return srv.ListenAndServe()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}