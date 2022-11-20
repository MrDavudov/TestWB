package handler

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/MrDavudov/TestWB/internal/repository"
	"github.com/MrDavudov/TestWB/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Server struct {
	httpServer 	*http.Server
}

func (s *Server) Start(port string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	s.httpServer = &http.Server{
		Addr:         	port,
		MaxHeaderBytes: 1<<20, // 1MB
		ReadTimeout:  	10 * time.Second,
		WriteTimeout: 	10 * time.Second,
	}

	// Инициализация config.yaml
	if err := initConfig(); err != nil {
		sugar.Fatalf("error Initializing configs: %s", err)
	}

	// Подключения .env
	if err := godotenv.Load(); err != nil {
		sugar.Fatalf("error loading env variables: %s", err)
	}

	// Подключение к БД
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

	// Взаимосвязи
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := NewHandler(port, services)

	// Проверка существует ли db.json, если нет то создать
	if err := FindJsonDB(); err != nil {
		sugar.Fatalf("failed create db.json: %s", err)
	}

	go func() {
		if err := handlers.services.SaveAsync(); err != nil {
			sugar.Fatalf("failed save async in db: %s", err)
		}
	}()

	// return srv.httpServer.ListenAndServe()
	handlers.router.Run(port)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func FindJsonDB() error {
	const jsonFile = "./db.json"

	_, err := os.Stat(jsonFile)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create("db.json")
			if err != nil {
				return err
			}

			err = os.WriteFile(jsonFile, []byte("[]"), 0666)
			if err != nil {
				return err
			}
		}
	}

	return nil
}