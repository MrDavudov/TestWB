package handler

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/MrDavudov/TestWB/internal/repository"
	"github.com/MrDavudov/TestWB/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	httpServer 	*http.Server
}

func (s *Server) Start(port string) error {
	logrus.SetFormatter(new(logrus.TextFormatter))

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
	if err := FindJsonDB(); err != nil {
		logrus.Fatalf("failed create db.json: %s", err)
	}

	// Взаимосвязи
	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := NewHandler(services)

	s.httpServer = &http.Server{
		Addr:         	port,
		Handler: handlers,
		MaxHeaderBytes: 1<<20, // 1MB
		ReadTimeout:  	10 * time.Second,
		WriteTimeout: 	10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
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