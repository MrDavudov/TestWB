package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MrDavudov/TestWB/internal/repository"
	"github.com/MrDavudov/TestWB/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	shutdownTimeout = 5 * time.Second
)

type Server struct {
	httpServer 	*http.Server
}

func (s *Server) Start(ctx context.Context) error {
	logrus.SetFormatter(new(logrus.TextFormatter))

	// Инициализация config.yaml
	if err := initConfig(); err != nil {
		logrus.Fatalf("error Initializing configs: %s", err)
	}

	// Подключения .env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err)
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
		logrus.Fatalf("failed to initialize db: %s", err)
	}

	// Проверка существует ли db.json, если нет то создать
	if err := repository.FindJsonDB(); err != nil {
		logrus.Fatalf("failed create db.json: %s", err)
	}

	// Взаимосвязи
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := NewHandler(services)

	s.httpServer = &http.Server{
		Addr:         	viper.GetString("port"),
		Handler: 		handlers,
		MaxHeaderBytes: 1<<20, // 1MB
		ReadTimeout:  	10 * time.Second,
		WriteTimeout: 	10 * time.Second,
	}

	// Запуск сервара
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen and serve: %v", err)
		}
	}()
	logrus.Infof("listening on localhost%s", s.httpServer.Addr)
	logrus.Info("Start server")

	// асинхронное обновление температуры каждые 60 секунд
	go func() {
		for {
			if err := services.SaveAsync(); err != nil {
				logrus.Fatalf("failed save async in db: %s", err)
			}
		}
	}()
	<-ctx.Done()

	// Плавное выход из программы
	logrus.Println("shutting down server gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown: %w", err) 
	}

	longShutdown := make(chan struct{}, 1)

	go func() {
		time.Sleep(3 * time.Second)
		longShutdown <- struct{}{}
	}()

	select {
	case <-shutdownCtx.Done():
		return fmt.Errorf("Server shutdown: %w", ctx.Err())
	case <-longShutdown:
		logrus.Println("Stop server")
	}

	return nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}